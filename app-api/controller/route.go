package controllers

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-xray-sdk-go/xray"
	stats "github.com/fukata/golang-stats-api-handler"
	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/aws"
	"github.com/supinf/supinf-mail/app-api/config"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/admin"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/basic"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/history"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/suppression"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
)

// Routes set API handlers
func Routes(api *operations.AppAPI) {
	// admin
	api.AdminPostAdminUsagePlanHandler = admin.PostAdminUsagePlanHandlerFunc(postUsagePlan)
	api.AdminPostAdminUsersHandler = admin.PostAdminUsersHandlerFunc(postUser)
	api.AdminPatchAdminUsersEnabledHandler = admin.PatchAdminUsersEnabledHandlerFunc(patchUserEnabled)

	// basic
	api.BasicPostMailsHandler = basic.PostMailsHandlerFunc(postMail)
	api.BasicPostBulkMailsHandler = basic.PostBulkMailsHandlerFunc(postBulkMail)

	// history
	api.HistoryGetMailsHistoriesHandler = history.GetMailsHistoriesHandlerFunc(listHistory)

	// suppression
	api.SuppressionGetSuppressionsHandler = suppression.GetSuppressionsHandlerFunc(listSuppression)
	api.SuppressionGetSuppressionsMailHandler = suppression.GetSuppressionsMailHandlerFunc(getSuppression)
	api.SuppressionPostSuppressionsHandler = suppression.PostSuppressionsHandlerFunc(postSuppression)
	api.SuppressionDeleteSuppressionsHandler = suppression.DeleteSuppressionsHandlerFunc(deleteSuppression)
}

//const (
//	accessFromCloudFront = "Amazon CloudFront"
//	accessFromPostman    = "PostmanRuntime"
//)

// Wrap wraps original HTTP handler
func Wrap(handler http.Handler) http.Handler {
	wrapped := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case eqauls(r, "/health"):
			w.WriteHeader(http.StatusOK)
			return
		case eqauls(r, "/version"):
			fmt.Fprintf(w, config.AppVersion)
			return
		case eqauls(r, "/api/v1/version"):
			fmt.Fprintf(w, config.AppVersion)
			return
		case eqauls(r, "/stats"):
			stats.Handler(w, r)
			return
		}

		// API 以外へのリクエストは弾く
		if !strings.HasPrefix(r.URL.Path, "/api/v") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx := auth.NewContext(r)

		started := misc.Now()
		ip, _ := aws.IPAddress(r)

		// 接続元が問題ないかを確認する
		if !isValidClient(r) {
			code := http.StatusForbidden
			accessLog(ctx, started, code, ip)
			w.WriteHeader(code)
			return
		}
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			r.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				if opErr, ok1 := err.(*net.OpError); ok1 {
					if sysErr, ok2 := opErr.Err.(*os.SyscallError); ok2 {
						if sysErr.Err == syscall.EPIPE {
							return
						}
					}
				}
				logs.Error(fmt.Sprint(err), nil, &logs.Map{
					"request-id": ctx.RequestID,
					"method":     ctx.Method,
					"path":       ctx.URL.Path,
					"address":    ip,
				})
				logs.StackTrace()
			}
		}()

		if config.AllowCORS {
			w.Header().Set("Access-Control-Allow-Origin", config.AllowCORSOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH,HEAD")
			w.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type,Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Type,Authorization,Date")
			w.Header().Set("Access-Control-Max-Age", "86400")
		}
		if strings.EqualFold(r.Method, http.MethodOptions) {
			w.WriteHeader(http.StatusOK)
			return
		}
		if config.SecuredTransport {
			if r.Header.Get("X-Forwarded-Proto") == "http" {
				http.Redirect(w, r, r.RequestURI, http.StatusMovedPermanently)
				w.Header().Add("Strict-Transport-Security", "max-age=31536000") // 1 year
				return
			}
		}

		ioWriter := w.(io.Writer)
		for _, val := range parseCsv(r.Header.Get("Accept-Encoding")) {
			if val == "gzip" {
				w.Header().Set("Content-Encoding", "gzip")
				g := gzip.NewWriter(w)
				defer g.Close()
				ioWriter = g
				break
			}
			if val == "deflate" {
				w.Header().Set("Content-Encoding", "deflate")
				z := zlib.NewWriter(w)
				defer z.Close()
				ioWriter = z
				break
			}
		}
		resp := &customResponseWriter{w, ioWriter, 200, 0}
		handler.ServeHTTP(resp, r)
		accessLog(ctx, started, resp.status, ip)
	})
	if config.AwsXRay {
		return xray.Handler(xray.NewFixedSegmentNamer(
			fmt.Sprintf("%s::%s", config.AppStage, config.ApplicationsName),
		), wrapped)
	}
	return wrapped
}

func isValidClient(r *http.Request) bool {
	return true
	//// ローカル環境であればどこからでも接続を許可
	//if !config.RunningOnAWS {
	//	return true
	//}
	//// CloudFront からの要求であれば許可
	//if strings.EqualFold(r.UserAgent(), accessFromCloudFront) {
	//	return true
	//}
	//// TravisCI (Newman での自動テスト & ELB 経由) からの要求であれば許可
	//if strings.HasPrefix(r.UserAgent(), accessFromPostman) {
	//	return true
	//}
	//return false
}

func accessLog(ctx *auth.Context, start time.Time, status int, address string) {
	if status == http.StatusNotFound {
		return
	}
	proc := fmt.Sprintf("%.3f", misc.Now().Sub(start).Seconds())
	logs.Info("ACCESS", nil, &logs.Map{
		"request-id": ctx.RequestID,
		"status":     status,
		"method":     ctx.Method,
		"path":       ctx.URL.Path,
		"address":    address,
		"proc":       proc,
	})
}

func eqauls(r *http.Request, url string) bool {
	return url == r.URL.Path
}

func parseCsv(data string) []string {
	splitted := strings.SplitN(data, ",", -1)
	parsed := make([]string, len(splitted))
	for i, val := range splitted {
		parsed[i] = strings.TrimSpace(val)
	}
	return parsed
}

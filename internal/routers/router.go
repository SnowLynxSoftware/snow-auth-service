package routers

import (
	"context"
	"net/http"
	"regexp"
	"snow-auth-service/internal/controllers"
	"strings"
)

var routes = []route{
	newRoute("GET", "/health", controllers.HealthController),
	newRoute("POST", "/auth/register", controllers.AuthControllerRegister),
	//newRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", apiUpdateWidget
	// Handles POST /api/widgets/([^/]+)/parts/([0-9]+)/update
	//func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request){
	//slug := getField(r, 0)
	//id, _ := strconv.Atoi(getField(r, 1))
	//fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
	//}
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

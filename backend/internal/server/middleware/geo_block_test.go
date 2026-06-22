//go:build unit

package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// newGeoBlockRouter 构造一个带 GeoBlock 中间件的测试路由，
// 平台路径 / 与 API 路径 /v1/messages 各注册一个 200 handler。
func newGeoBlockRouter(cfg config.GeoBlockConfig) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(GeoBlock(cfg))
	ok := func(c *gin.Context) { c.String(http.StatusOK, "ok") }
	r.GET("/", ok)
	r.POST("/v1/messages", ok)
	r.GET("/health", ok)
	r.GET("/api/v1/user/profile", ok)
	return r
}

func doGeoReq(r *gin.Engine, method, path, country, accept string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	if country != "" {
		req.Header.Set("CF-IPCountry", country)
	}
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGeoBlock_DisabledAllowsAll(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          false,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "CN", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_EmptyListAllowsAll(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"", "  "},
	})
	w := doGeoReq(r, http.MethodGet, "/", "CN", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_BlockedCountryOnPlatform(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "CN", "")
	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestGeoBlock_BlockedCountryOnPlatformAPI(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	// /api/v1/* 属于平台接口，应被拦截。
	w := doGeoReq(r, http.MethodGet, "/api/v1/user/profile", "CN", "")
	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestGeoBlock_AllowedCountry(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "US", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_BlockedCountryAPIGatewayBypass(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	// API 网关流量即便来自被封国家也放行。
	w := doGeoReq(r, http.MethodPost, "/v1/messages", "CN", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_HealthBypass(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/health", "CN", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_MissingHeaderFailsOpen(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "", "")
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGeoBlock_UnknownCountryFailsOpen(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	for _, code := range []string{"XX", "T1"} {
		w := doGeoReq(r, http.MethodGet, "/", code, "")
		require.Equalf(t, http.StatusOK, w.Code, "country %s should fail open", code)
	}
}

func TestGeoBlock_CaseInsensitive(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"cn"}, // 配置小写
	})
	w := doGeoReq(r, http.MethodGet, "/", "cn", "") // Header 小写
	require.Equal(t, http.StatusForbidden, w.Code)
}

func TestGeoBlock_HTMLResponseForBrowser(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "CN", "text/html,application/xhtml+xml")
	require.Equal(t, http.StatusForbidden, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "text/html")
	require.Contains(t, w.Body.String(), "<!DOCTYPE html>")
}

func TestGeoBlock_JSONResponseForNonBrowser(t *testing.T) {
	r := newGeoBlockRouter(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
	})
	w := doGeoReq(r, http.MethodGet, "/", "CN", "application/json")
	require.Equal(t, http.StatusForbidden, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "application/json")
	require.Contains(t, w.Body.String(), "region_not_supported")
	require.False(t, strings.Contains(w.Body.String(), "<!DOCTYPE"))
}

func TestGeoBlock_CustomHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(GeoBlock(config.GeoBlockConfig{
		Enabled:          true,
		BlockedCountries: []string{"CN"},
		CountryHeader:    "X-Country-Code",
	}))
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Country-Code", "CN")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusForbidden, w.Code)
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/Wei-Shaw/sub2api/internal/config"
	"github.com/gin-gonic/gin"
)

// geoBlockAPIPrefixes 是 API 网关路径前缀，命中则跳过地区封锁。
// 这些是面向程序的 API 流量（Claude/OpenAI/Gemini 兼容层），按需求“API 访问暂不限制”。
// 与 internal/web.shouldBypassEmbeddedFrontend 的分类保持一致。
var geoBlockAPIPrefixes = []string{
	"/v1/",
	"/v1beta/",
	"/backend-api/",
	"/antigravity/",
	"/responses/",
	"/images/",
}

// geoBlockAPIExact 是需放行的精确 API 路径（无尾部 path 段的别名）。
var geoBlockAPIExact = map[string]struct{}{
	"/health":                  {},
	"/api/event_logging/batch": {}, // Claude Code 遥测
	"/responses":               {},
	"/chat/completions":        {},
	"/embeddings":              {},
}

// isGeoBlockExemptPath 判断路径是否属于 API 网关流量（不受地区封锁影响）。
func isGeoBlockExemptPath(path string) bool {
	if _, ok := geoBlockAPIExact[path]; ok {
		return true
	}
	for _, p := range geoBlockAPIPrefixes {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

// GeoBlock 返回按来源国家/地区限制平台访问的中间件。
//
// 仅限制平台访问（前端页面 + /api/v1 管理/认证/支付接口）；API 网关流量始终放行。
// 国家码从 cfg.CountryHeader（默认 CF-IPCountry，由 Cloudflare 注入）读取，
// 命中 cfg.BlockedCountries 则返回 403。Header 缺失或值未知（如 XX/T1）按 fail-open 放行。
func GeoBlock(cfg config.GeoBlockConfig) gin.HandlerFunc {
	if !cfg.Enabled {
		return func(c *gin.Context) { c.Next() }
	}

	header := strings.TrimSpace(cfg.CountryHeader)
	if header == "" {
		header = "CF-IPCountry"
	}

	// 预构建大写国家码集合，热路径零分配匹配。
	blocked := make(map[string]struct{}, len(cfg.BlockedCountries))
	for _, code := range cfg.BlockedCountries {
		normalized := strings.ToUpper(strings.TrimSpace(code))
		if normalized != "" {
			blocked[normalized] = struct{}{}
		}
	}

	// 没有配置任何国家码 → 等效关闭，避免误伤。
	if len(blocked) == 0 {
		return func(c *gin.Context) { c.Next() }
	}

	return func(c *gin.Context) {
		// API 网关流量放行。
		if isGeoBlockExemptPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		country := strings.ToUpper(strings.TrimSpace(c.GetHeader(header)))
		if country == "" {
			// Header 缺失（绕过 Cloudflare 直连源站等）→ fail-open。
			c.Next()
			return
		}

		if _, hit := blocked[country]; !hit {
			c.Next()
			return
		}

		writeGeoBlockResponse(c)
		c.Abort()
	}
}

// writeGeoBlockResponse 按 Accept 头协商返回 403：浏览器返回 HTML 提示页，其余返回 JSON。
func writeGeoBlockResponse(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept"), "text/html") {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusForbidden, geoBlockHTML)
		return
	}
	c.JSON(http.StatusForbidden, gin.H{
		"code":    "region_not_supported",
		"message": "Access from your region is not supported. / 当前地区暂不支持访问。",
	})
}

// geoBlockHTML 是被封锁地区访问平台页面时返回的中英双语提示页。
const geoBlockHTML = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>暂不支持访问 · Region Not Supported</title>
<style>
  body { margin:0; min-height:100vh; display:flex; align-items:center; justify-content:center;
    font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"PingFang SC","Microsoft YaHei",sans-serif;
    background:#0b0f17; color:#e6edf3; }
  .card { max-width:520px; padding:48px 40px; text-align:center; }
  .code { font-size:13px; letter-spacing:2px; color:#7d8590; text-transform:uppercase; }
  h1 { font-size:24px; margin:16px 0 8px; font-weight:600; }
  p { font-size:15px; line-height:1.7; color:#9aa4b2; margin:8px 0; }
  .en { color:#6e7681; font-size:14px; }
</style>
</head>
<body>
  <div class="card">
    <div class="code">403 · Region Not Supported</div>
    <h1>当前地区暂不支持访问</h1>
    <p>抱歉，本平台暂未向你所在的国家或地区开放。</p>
    <p class="en">Sorry, this platform is not currently available in your country or region.</p>
  </div>
</body>
</html>`

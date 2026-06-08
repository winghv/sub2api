export const DEFAULT_SITE_NAME = 'Superwhv'
export const DEFAULT_SITE_SUBTITLE = '助力碳基进化'
export const DEFAULT_SITE_SUBTITLE_EN = 'Empowering Carbon-Based Evolution'
export const LEGACY_MIXED_SITE_SUBTITLE = `${DEFAULT_SITE_SUBTITLE} · ${DEFAULT_SITE_SUBTITLE_EN}`
export const LOCALIZED_DEFAULT_SITE_SUBTITLES = [
  DEFAULT_SITE_SUBTITLE,
  DEFAULT_SITE_SUBTITLE_EN,
  LEGACY_MIXED_SITE_SUBTITLE
]
export const DEFAULT_SITE_LOGO = '/superwhv-logo.png'
export const DEFAULT_DOCUMENT_TITLE = `${DEFAULT_SITE_NAME} - AI API Gateway`

export function isDefaultSiteSubtitle(value?: string | null): boolean {
  const subtitle = value?.trim()
  return !subtitle || LOCALIZED_DEFAULT_SITE_SUBTITLES.includes(subtitle)
}

/**
 * Shared theme helpers — single source of truth for the light/dark default.
 *
 * The `dark` class on `<html>` drives all Tailwind dark: variants and the
 * cyber shells (cyber-home / app-shell / auth-shell pick their `--dark` or
 * `--light` modifier from it).
 *
 * Fork default: first-time visitors (no saved preference) get **dark** (the
 * cyber look) instead of following `prefers-color-scheme`. An explicit choice
 * is always respected.
 */

const THEME_STORAGE_KEY = 'theme'

/** Resolve the initial theme. Only an explicit saved 'light' keeps light mode. */
export function resolveInitialTheme(): boolean {
  return localStorage.getItem(THEME_STORAGE_KEY) !== 'light'
}

/** Apply the theme class to <html> without persisting. */
export function applyThemeClass(dark: boolean): void {
  document.documentElement.classList.toggle('dark', dark)
}

/** Persist the user's explicit theme choice. */
export function persistTheme(dark: boolean): void {
  localStorage.setItem(THEME_STORAGE_KEY, dark ? 'dark' : 'light')
}

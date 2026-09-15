const PLACEHOLDER = /\{\{\s*([^}]+?)\s*\}\}/g;

/**
 * Replaces `{{ env.NAME }}` placeholders with values from `env`. Used to interpolate
 * env vars into the raw testflowkit.yml text before it is parsed, mirroring the Go loader.
 */
export function substituteEnv(text: string, env: Record<string, string>): string {
  return text.replace(PLACEHOLDER, (match, expr: string) => {
    if (!expr.startsWith('env.')) {
      return match;
    }
    const name = expr.slice('env.'.length).trim();
    return env[name] ?? match;
  });
}

/**
 * Replaces `{{ NAME }}` (scenario variable) and `{{ env.NAME }}` placeholders in a step
 * argument string. Unknown placeholders are left untouched.
 */
export function substituteVariables(
  text: string,
  variables: Map<string, string>,
  env: Record<string, string> = process.env as Record<string, string>,
): string {
  return text.replace(PLACEHOLDER, (match, expr: string) => {
    if (expr.startsWith('env.')) {
      const name = expr.slice('env.'.length).trim();
      return env[name] ?? match;
    }
    return variables.get(expr) ?? match;
  });
}

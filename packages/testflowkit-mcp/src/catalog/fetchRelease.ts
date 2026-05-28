const FETCH_TIMEOUT_MS = 15_000;

export async function fetchUrl(url: string): Promise<string> {
  const res = await fetch(url, {
    signal: AbortSignal.timeout(FETCH_TIMEOUT_MS),
  });
  if (!res.ok) {
    throw new Error(`HTTP ${res.status} fetching ${url}`);
  }
  return res.text();
}

export async function fetchReleaseAsset(
  repository: string,
  tag: string,
  asset: string,
): Promise<string> {
  const url = `https://github.com/${repository}/releases/download/${tag}/${asset}`;
  return fetchUrl(url);
}

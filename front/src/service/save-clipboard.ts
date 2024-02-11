export async function saveClipboard(value: string) {
  await navigator.clipboard.writeText(value);
}

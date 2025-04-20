// place files you want to import through the `$lib` alias in this folder.
// Format file size
export function formatFileSize(bytes) {
  if (bytes < 1024) return bytes + " B";
  else if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(2) + " KB";
  else return (bytes / (1024 * 1024)).toFixed(2) + " MB";
}

// Format date
export function formatDate(dateString) {
  const date = new Date(dateString);
  return date.toLocaleString();
}

export const useMedia = () => {
  const getMediaUrl = (path?: string | null): string => {
    if (!path) return ''
    if (path.startsWith('http://') || path.startsWith('https://') || path.startsWith('data:')) {
      return path
    }
    if (path.startsWith('/uploads')) {
      return path
    }
    if (path.startsWith('uploads/')) {
      return `/${path}`
    }
    return `/uploads/${path}`
  }

  const getCompanyInitial = (name?: string | null): string => {
    if (!name) return 'B'
    const cleanName = name.trim().replace(/^(Công ty TNHH|Công ty CP|Công ty Cổ phần|Tập đoàn|Doanh nghiệp)\s+/i, '')
    return (cleanName.charAt(0) || name.charAt(0) || 'B').toUpperCase()
  }

  return {
    getMediaUrl,
    getCompanyInitial
  }
}

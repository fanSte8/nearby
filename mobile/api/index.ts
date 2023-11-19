export const formatError = (error: any) => {
  if (typeof error.response.data.error === 'object') {
    return Object.values(error.response.data.error).join('\n')
  }

  return error.response.data.error
}
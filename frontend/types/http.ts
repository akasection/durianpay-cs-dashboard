export type BaseResponse<T extends object = Record<string, any>> = {
  code: number
  message: string
  data: T
  meta?: {
    total?: number
    offset?: number
    limit?: number
    extra?: Record<string, any>
  }
}

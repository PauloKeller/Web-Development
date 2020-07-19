import Axios, { AxiosInstance, AxiosRequestConfig } from 'axios'

export default class Client {
  baseURL: string
  instance: AxiosInstance
  headers: any = {
    'Accept-language': 'pt-br',
    'Content-Type': 'application/json',
  }

  constructor(baseURL: string, options?: AxiosRequestConfig, headers?: any) {
    this.baseURL = baseURL

    if (headers) {
      this.headers = headers
    }

    this.instance = Axios.create({
      baseURL: this.baseURL,
      headers: this.headers,
      ...options,
    })
  }
}

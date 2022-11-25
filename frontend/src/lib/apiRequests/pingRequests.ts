import { AxiosResponse } from 'axios'
import { getAsync } from './api'

type Ping = {
  message: string
}

export const getPing = async (): Promise<AxiosResponse<Ping>> => getAsync<Ping>(`/ping`)

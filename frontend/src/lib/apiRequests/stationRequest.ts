import { AxiosResponse } from 'axios'
import { getAsync } from './api'

export type Station = {
  fid: number
  id: number
  name: string
  nameSwedish: string
  nameEnglish: string
  address: string
  addressSwedish: string
  city: string
  citySwedish: string
  operator: string
  capacity: number
  coordinates: { x: number; y: number }
}

export const getStations = async (page: number, str: string): Promise<AxiosResponse<Station[]>> =>
  str === ''
    ? getAsync<Station[]>(`/api/stations/page/${page}`)
    : getAsync<Station[]>(`/api/stations/page/${page}/${str}`)

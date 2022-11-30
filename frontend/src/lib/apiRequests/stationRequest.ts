import { AxiosResponse } from 'axios'
import { getAsync } from './api'

type Station = {
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

export const getStations = async (page: number): Promise<AxiosResponse<Station[]>> =>
  getAsync<Station[]>(`/api/stations/page/${page}`)

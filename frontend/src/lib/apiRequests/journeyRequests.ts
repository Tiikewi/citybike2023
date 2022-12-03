import { AxiosResponse } from 'axios'
import { getAsync } from './api'

type Journey = {
  id: number
  departureTime: string
  returnTime: string
  departureStationId: number
  departureStationName: string
  returnStationId: number
  returnStationName: string
  distance: number
  duration: number
}

export const getJourneys = async (page: number, sort: number): Promise<AxiosResponse<Journey[]>> =>
  getAsync<Journey[]>(`/api/journeys/page/${page}/${sort}`)

import { PositionsDto } from '@/hub';
import { apiInstance } from '@/services/api-base';

export const GetTrail = async (vehicleId: string): Promise<PositionsDto> => {
  const res = await apiInstance.get(`trail/${vehicleId}`);
  return res?.data as PositionsDto;
};

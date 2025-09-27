import { RegionList } from '@/types/TableInterfaces';
import axios from 'axios';

const baseUrl = import.meta.env.VITE_BASE_URL;

axios.defaults.baseURL = baseUrl;

export async function listRegionPlate(): Promise<RegionList[]> {
  try {
    const data = await axios.get(`plate-code/region`, {
      headers: {
        'Content-Type': 'application/json',
      },
    });
    console.log(data.data.data);

    return data.data.data;
  } catch (error) {
    console.log(error);
    throw error;
  }
}

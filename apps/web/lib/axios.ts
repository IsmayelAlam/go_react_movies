/* eslint-disable @typescript-eslint/no-explicit-any */
import axios, { RawAxiosRequestHeaders } from 'axios';

export interface QueryProps {
  endpoint?: string;
  method?: 'get' | 'post' | 'put' | 'delete';
  disable?: boolean;

  data?: any;
  params?: {
    filters?: any;
    populate?: any;
    sort?: any;
    pagination?: { page: number; pageSize: number };
  } & Record<string, any>;
  id?: string | number;
  headers?: Partial<RawAxiosRequestHeaders>;
}

export async function axiosCall(props: QueryProps) {
  try {
    const {
      method = 'get',
      data,
      endpoint,
      id,
      params,
      headers,
      disable,
    } = props;

    if (!endpoint || !method) throw new Error('Endpoint/Method is missing');

    const url = `http://localhost:8080/${endpoint}${id ? '/' + id : ''}`;

    const config = {
      method,
      maxBodyLength: Infinity,
      url,
      headers,
      params,
      data,
    };

    if (disable) return { data: null, meta: null };

    const { data: response } = await axios.request(config);
    return response;
  } catch (error: any) {
    throw error.response.data.error;
  }
}

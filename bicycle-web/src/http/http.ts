import axios from 'axios';

export const api = axios.create({
  baseURL: 'https://bicycle-api-dot-strangersfeir.uc.r.appspot.com/v1/bicycle',
});
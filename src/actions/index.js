import { API_BASE_URL } from '../config';
import Cookies from 'js-cookie';
export default function apiFetch(endpoint, options = { headers: {} }) {
  let token = getLocally('jwt');
  if (token) options.headers['Authorization'] = `Bearer ${token}`;

  return fetch(`${API_BASE_URL}/${endpoint}`, options);
}

export function persistLocally(key, value) {
  Cookies.set(key, JSON.stringify(value));
}
export function getLocally(key) {
  let data = Cookies.get(key);
  return data === undefined ? undefined : JSON.parse(data);
}

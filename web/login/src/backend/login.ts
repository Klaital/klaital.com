

export interface User {
  id: number;
  username: string;
  email: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface RegisterResponse {
  sessionToken: string;
  user: User | undefined;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  sessionToken: string;
  user: User | undefined;
}

export interface LogoutRequest {
  sessionToken: string;
}

export interface LogoutResponse {
}


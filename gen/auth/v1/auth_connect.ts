// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file auth/v1/auth.proto (package auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { LoginRequest, LoginResponse } from "./auth_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service auth.v1.AuthService
 */
export const AuthService = {
  typeName: "auth.v1.AuthService",
  methods: {
    /**
     * @generated from rpc auth.v1.AuthService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;


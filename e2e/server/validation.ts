import Ajv from "ajv";
import addFormats from "ajv-formats";
import type { Request, Response, NextFunction } from "express";

const ajv = new Ajv({ allErrors: true });
addFormats(ajv);

// Type definitions
export interface Post {
  id?: number;
  title: string;
  body: string;
  userId: number;
}

export interface Album {
  id?: number;
  title: string;
  userId: number;
}

export interface Photo {
  id?: number;
  title: string;
  albumId: number;
  url: string;
  thumbnailUrl: string;
}

export interface User {
  id?: number;
  name: string;
  username: string;
  email: string;
  phone?: string;
  website?: string;
}

export interface ValidationError {
  keyword: string;
  dataPath: string;
  schemaPath: string;
  params: Record<string, any>;
  message: string;
}

// JSON Schema definitions for validation
export const schemas = {
  post: {
    type: "object",
    required: ["title", "body", "userId"],
    properties: {
      title: { type: "string", minLength: 1, maxLength: 200 },
      body: { type: "string", minLength: 1, maxLength: 1000 },
      userId: { type: "integer", minimum: 1 },
    },
    additionalProperties: false,
  },

  album: {
    type: "object",
    required: ["title", "userId"],
    properties: {
      id: { type: "integer", minimum: 1 },
      title: { type: "string", minLength: 1, maxLength: 100 },
      userId: { type: "integer", minimum: 1 },
    },
    additionalProperties: false,
  },

  photo: {
    type: "object",
    required: ["title", "albumId", "url", "thumbnailUrl"],
    properties: {
      id: { type: "integer", minimum: 1 },
      title: { type: "string", minLength: 1, maxLength: 100 },
      albumId: { type: "integer", minimum: 1 },
      url: { type: "string", format: "uri" },
      thumbnailUrl: { type: "string", format: "uri" },
    },
    additionalProperties: false,
  },

  user: {
    type: "object",
    required: ["name", "username", "email"],
    properties: {
      id: { type: "integer", minimum: 1 },
      name: { type: "string", minLength: 1, maxLength: 50 },
      username: { type: "string", minLength: 1, maxLength: 50 },
      email: { type: "string", format: "email" },
      phone: { type: "string" },
      website: { type: "string", format: "uri" },
    },
    additionalProperties: false,
  },
};

// Validation middleware
export const validationMiddleware = (
  req: Request,
  res: Response,
  next: NextFunction,
): void => {
  const path = req.path;
  const method = req.method;

  // Skip validation for GET and DELETE requests
  if (method === "GET" || method === "DELETE") {
    return next();
  }

  // Define validation rules: array of methods + path combination
  const validationRules: Array<{
    methods: string[];
    pathPattern: string;
    schema: any;
  }> = [
    {
      methods: ["POST", "PUT", "PATCH"],
      pathPattern: "/posts",
      schema: schemas.post,
    },
    {
      methods: ["POST", "PUT", "PATCH"],
      pathPattern: "/albums",
      schema: schemas.album,
    },
    {
      methods: ["POST", "PUT", "PATCH"],
      pathPattern: "/photos",
      schema: schemas.photo,
    },
    {
      methods: ["POST", "PUT", "PATCH"],
      pathPattern: "/users",
      schema: schemas.user,
    },
  ];

  // Find matching validation rule
  const rule = validationRules.find(
    (rule) => rule.methods.includes(method) && path.includes(rule.pathPattern),
  );

  if (!rule) {
    return next();
  }

  // Validate request body against the schema
  const validate = ajv.compile(rule.schema);
  const valid = validate(req.body);

  if (!valid) {
    res.status(400).json({
      error: "Validation failed",
      details: validate.errors,
      message: `Request data does not match required schema for ${method} ${path}`,
    });
    return;
  }

  // Add custom business logic validation
  if (path.includes("/posts") && req.body.userId) {
    // Check if user exists (in a real app, this would query the database)
    if (req.body.userId > 10) {
      res.status(400).json({
        error: "Business validation failed",
        message: "User ID does not exist",
      });
      return;
    }
  }

  if (path.includes("/photos") && req.body.albumId) {
    // Check if album exists
    if (req.body.albumId > 10) {
      res.status(400).json({
        error: "Business validation failed",
        message: "Album ID does not exist",
      });
      return;
    }
  }

  next();
};

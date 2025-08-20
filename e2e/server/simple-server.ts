import jsonServer from "json-server";
import express from "express";
import type { Request, Response } from "express";
import { validationMiddleware, schemas } from "./validation.ts";
import Ajv from "ajv";
import addFormats from "ajv-formats";

const server = jsonServer.create();
const router = jsonServer.router("./db.json");
const middlewares = jsonServer.defaults();

server.use(middlewares);

server.use(express.json());
server.use(express.urlencoded({ extended: true }));

server.use(validationMiddleware);

server.get("/posts/:id", (req: Request, res: Response) => {
  const postId = parseInt(req.params.id);

  const mockPosts = [
    {
      id: 1,
      userId: 1,
      title:
        "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
      body: "quia et suscipit suscipit recusandae consequuntur expedita et cum reprehenderit molestiae ut ut quas totam nostrum rerum est autem sunt rem eveniet architecto",
    },
    {
      id: 2,
      userId: 1,
      title: "qui est esse",
      body: "est rerum tempore vitae sequi sint nihil reprehenderit dolor beatae ea dolores neque fugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis qui aperiam non debitis possimus qui neque nisi nulla",
    },
  ];

  const post = mockPosts.find((p) => p.id === postId);

  if (post) {
    res.json(post);
  } else {
    res.status(404).json({
      error: "Post not found",
      message: `Post with ID ${postId} does not exist`,
    });
  }
});

server.post("/posts", (req: Request, res: Response) => {
  const newPost = {
    id: Date.now(),
    ...req.body,
  };

  res.status(201).json(newPost);
});

server.delete("/posts/:id", (req: Request, res: Response) => {
  const postId = parseInt(req.params.id);

  res.status(200).json({
    success: true,
    message: `Post with ID ${postId} has been deleted successfully`,
    deletedPostId: postId,
  });
});

server.get("/albums/:id", (req: Request, res: Response) => {
  const albumId = parseInt(req.params.id);

  const mockAlbums = [
    {
      id: 1,
      userId: 1,
      title: "quidem molestiae enim",
    },
    {
      id: 2,
      userId: 1,
      title: "sunt qui excepturi placeat culpa",
    },
  ];

  const album = mockAlbums.find((a) => a.id === albumId);

  if (album) {
    res.json(album);
  } else {
    res.status(404).json({
      error: "Album not found",
      message: `Album with ID ${albumId} does not exist`,
    });
  }
});

server.get("/albums/:id/photos", (req: Request, res: Response) => {
  const albumId = parseInt(req.params.id);

  const mockPhotos = [
    {
      id: 1,
      albumId: 1,
      title: "accusamus beatae ad facilis cum similique qui sunt",
      url: "https://via.placeholder.com/600/92c952",
      thumbnailUrl: "https://via.placeholder.com/150/92c952",
    },
    {
      id: 2,
      albumId: 1,
      title: "reprehenderit est deserunt velit ipsam",
      url: "https://via.placeholder.com/600/771796",
      thumbnailUrl: "https://via.placeholder.com/150/771796",
    },
  ];

  const photos = mockPhotos.filter((p) => p.albumId === albumId);

  if (photos.length > 0) {
    res.json(photos);
  } else {
    res.status(404).json({
      error: "Photos not found",
      message: `No photos found for album ID ${albumId}`,
    });
  }
});

server.get("/echo", (req: Request, res: Response) => {
  res.jsonp(req.query);
});

server.post("/validate", (req: Request, res: Response) => {
  const { data, type } = req.body;

  if (!data || !type) {
    return res.status(400).json({
      error: "Missing required fields",
      message: 'Both "data" and "type" fields are required',
    });
  }

  const availableSchemas = {
    post: schemas.post,
    album: schemas.album,
    photo: schemas.photo,
    user: schemas.user,
  };

  const schema = availableSchemas[type as keyof typeof availableSchemas];
  if (!schema) {
    return res.status(400).json({
      error: "Invalid type",
      message: `Type must be one of: ${Object.keys(availableSchemas).join(
        ", ",
      )}`,
    });
  }

  const ajv = new Ajv({ allErrors: true });
  addFormats(ajv);

  const validate = ajv.compile(schema);
  const valid = validate(data);

  if (valid) {
    res.json({
      valid: true,
      message: "Data is valid according to schema",
    });
  } else {
    res.status(400).json({
      valid: false,
      errors: validate.errors,
      message: "Data validation failed",
    });
  }
});

server.get("/health", (req: Request, res: Response) => {
  res.json({
    status: "healthy",
    timestamp: new Date().toISOString(),
    service: "TestFlowKit Mock Server",
    version: "1.0.0",
  });
});

server.use(router);

const port = process.env.PORT || 3001;
server.listen(port, () => {
  console.log(`🚀 TestFlowKit Mock Server is running on port ${port}`);
  console.log(`📊 API endpoints available at http://localhost:${port}`);
  console.log(`🔍 Health check: http://localhost:${port}/health`);
});

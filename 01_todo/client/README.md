# Todo UI - React + Vite + Bun + Tailwind

A simple Todo application UI built with React, Vite, Bun, and Tailwind CSS.

## Prerequisites

- Bun installed on your system
- The Go backend API running on port 5001

## Setup

1. Install dependencies:

```bash
bun install
```

2. Start the development server:

```bash
bun run dev
```

The app will be available at `http://localhost:3000`

## Features

- Add new todos with title and date
- Mark todos as done
- Delete todos
- Clean, minimal UI with Tailwind CSS

## API Configuration

The app expects the backend API to be running at `http://localhost:5001/todos`

If your API is running on a different port, update the `API_URL` in `src/App.jsx`

## Build for Production

```bash
bun run build
```

## Preview Production Build

```bash
bun run preview
```

"use client";

import Image from "next/image";
import { Inter } from "next/font/google";
import styles from "./page.module.css";
import React from "react";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });
export type Todo = {
  id: number;
  title: string;
  category: string | null;
  description: string | null;
  createdAt: string;
  updatedAt: string;
  deletedAt: string | null;
};
export default function Home() {
  const onSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const data = Object.fromEntries(
      Array.from(formData.entries()).map(([key, value]) => [
        key,
        value === "" ? null : value,
      ]),
    );
    fetch("http://localhost:8080/post", {
      method: "POST",
      body: JSON.stringify(data),
    })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  const [todoList, setTodoList] = React.useState<Todo[]>([]);
  React.useEffect(() => {
    const fetchTodoList = async () => {
      try {
        const response = await fetch("http://localhost:8080/posts");
        const data = await response.json();
        setTodoList(data.data || []);
        return;
      } catch (e) {
        console.error(e);
      }
    };
    fetchTodoList();
  }, []);
  return (
    <main>
      <div style={{ maxWidth: "25vw", margin: "0 auto" }}>
        <div>
          <h1>TODO LIST</h1>
          <div>
            <ul>
              {todoList.map((todo) => (
                <li key={todo.id}>
                  <Link href={`/post/${todo.id}`}>{todo.title}</Link>
                </li>
              ))}
            </ul>
          </div>
        </div>

        <form
          style={{ display: "flex", flexDirection: "column" }}
          onSubmit={onSubmit}
        >
          <label htmlFor="title">Title</label>
          <input type="text" id="title" name="Title" />
          <label htmlFor="category">Category</label>
          <input type="category" id="category" name="category" />
          <label htmlFor="description">Description</label>
          <textarea id="description" name="description" />
          <button type="submit" style={{ marginTop: "16px" }}>
            Send
          </button>
        </form>
      </div>
    </main>
  );
}

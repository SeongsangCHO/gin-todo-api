"use client";

import React from "react";
import { Todo } from "@/app/page";
import { useRouter } from "next/navigation";

export default function PostDetail({ params }: { params: { id: string } }) {
  const router = useRouter();
  console.log(params.id);
  const [todo, setTodo] = React.useState<Todo | null>(null);
  React.useEffect(() => {
    async function fetchPost() {
      const response = await fetch(`http://localhost:8080/post/${params.id}`);
      const data = await response.json();
      setTodo(data);
    }
    fetchPost();
  }, []);
  console.log(todo);

  const onDelete = async () => {
    const response = await fetch(`http://localhost:8080/post/${params.id}`, {
      method: "DELETE",
    });
    if (response.status === 200) {
      router.push("/");
    }
  };
  if (!todo) return <div>Loading..</div>;
  return (
    <div>
      <button onClick={onDelete}>DELETE</button>
      <div style={{ display: "flex", flexDirection: "column", gap: "16px" }}>
        <h1>Post Detail</h1>
        <span>{todo?.id}</span>
        <span>{todo?.title}</span>
        <span>{todo?.category}</span>
        <span>{todo?.description}</span>
        <span>{todo?.createdAt}</span>
        <span>{todo?.updatedAt}</span>
        <span>{todo?.deletedAt}</span>
      </div>
    </div>
  );
}

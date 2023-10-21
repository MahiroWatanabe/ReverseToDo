import React from "react";
import styles from "./TodoItem.module.scss";
import { useRouter } from "next/router";

type Props = {
  item: {
    CreatedAt: string;
    DeletedAt: string | null;
    ID: number;
    UpdatedAt: string;
    assignid: number;
    createid: number;
    deadline: string;
    description: string;
    status: number;
    title: string;
  };
};

const TodoItem = ({ item }: Props) => {
  const router = useRouter();
  const currentPath = router.asPath;

  const handleSubmit = () => {
    router.push(`${currentPath}/task/${item.ID}`);
  };
  return (
    <ul className={styles["todo-ui"]} onClick={handleSubmit}>
      <li key={item.ID}>
        <h3>{item.title}</h3>
        <p>{item.description}</p>
        <p>Deadline: {item.deadline}</p>
        <p>Status: {item.status}</p>
      </li>
    </ul>
  );
};

export default TodoItem;

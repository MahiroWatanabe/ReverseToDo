import React, { useState } from "react";
import styles from "./UpdateSide.module.scss";

type SidebarProps = {
  useTaskData: ReturnTaskData | null;
  handleUpdateTodo: (
    newTitle: any,
    newDescription: any,
    newDeadline: any
  ) => void;
};

const UpdataSide: React.FC<SidebarProps> = ({
  useTaskData,
  handleUpdateTodo,
}) => {
  const [newTitle, setNewTitle] = useState(useTaskData?.title);
  const [newDescription, setNewDescription] = useState(
    useTaskData?.description
  );
  const [newDeadline, setNewDeadline] = useState(useTaskData?.deadline);

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewTitle(e.target.value);
  };

  const handleDescriptionChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewDescription(e.target.value);
  };

  const handleDeadlineChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewDeadline(e.target.value);
  };

  const handleUpdateClick = () => {
    handleUpdateTodo(newTitle, newDescription, newDeadline);
  };

  return (
    <div className={styles["search-container"]}>
      <h2>Update Task</h2>
      <label>
        Title:
        <input type="text" value={newTitle} onChange={handleTitleChange} />
      </label>
      <label>
        Description:
        <input
          type="text"
          value={newDescription}
          onChange={handleDescriptionChange}
        />
      </label>
      <label>
        Deadline:
        <input
          type="text"
          value={newDeadline}
          onChange={handleDeadlineChange}
        />
      </label>
      <button onClick={handleUpdateClick}>Update Task</button>
    </div>
  );
};

export default UpdataSide;

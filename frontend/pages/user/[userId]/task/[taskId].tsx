import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import {
  getUserTaskData,
  putTaskStatus,
  putTaskData,
} from "../../../../utils/userDataRequest";
import styles from "./task.module.scss";
import UpdataSide from "../../../../components/UpdateSide/UpdataSide";

const TaskComponent = () => {
  const [useTaskData, setUseTaskData] = useState<ReturnTaskData | null>(null);
  const [isChecked, setIsChecked] = useState(useTaskData?.status == 1);
  const [openSide, setOpenSide] = useState(false);

  const router = useRouter();
  const currentPath: string = router.asPath;
  const pathParts = currentPath.split("/");

  const handleUpdate = () => {
    setOpenSide((pre) => !pre);
  };

  const handlegetTask = (url: string) => {
    getUserTaskData(url).then((res) => {
      setUseTaskData(res);
      return res;
    });
  };

  const handleupdateStatus = (url: string, flag: number, id: any) => {
    putTaskStatus(url, flag, id).then((res) => {
      setUseTaskData(res);
      return res;
    });
  };

  const handleUpdateTask = (url: string, putTask: ReturnTaskData | null) => {
    putTaskData(url, putTask).then((res) => {
      return res;
    });
  };

  const handleUpdateTodo = (
    newTitle: string,
    newDescription: string,
    newDeadline: string
  ) => {
    if (useTaskData) {
      setUseTaskData({
        ...useTaskData,
        title: newTitle,
        description: newDescription,
        deadline: newDeadline,
      });
    }
  };

  const handleDelete = () => {};

  const handleCheckChange = () => {
    setIsChecked((pre) => !pre);
    handleupdateStatus(`/task`, isChecked ? 1 : 0, useTaskData?.ID);
  };

  useEffect(() => {
    handlegetTask(`/task?userid=${pathParts[2]}&taskid=${pathParts[4]}`);
  }, [isChecked]);

  useEffect(() => {
    handleUpdateTask(`/task`, useTaskData);
  }, [useTaskData]);

  return (
    <div className={styles["parent-container"]}>
      <div className={styles["user-task-container"]}>
        <h1 className={styles.title}>{useTaskData?.title}</h1>
        <p className={styles.description}>{useTaskData?.description}</p>
        <p className={styles.status}>
          Status: {isChecked ? "Active" : "Inactive"}
        </p>
        <p className={styles.deadline}>Deadline: {useTaskData?.deadline}</p>
        <button onClick={handleUpdate}>Update</button>
        <button onClick={handleDelete}>Delete</button>
        <label>
          <input
            type="checkbox"
            className={styles.checkbox}
            checked={isChecked}
            onChange={handleCheckChange}
          />
          Mark as Active
        </label>
      </div>
      {openSide && (
        <UpdataSide
          useTaskData={useTaskData}
          handleUpdateTodo={handleUpdateTodo}
        />
      )}
    </div>
  );
};

export default TaskComponent;

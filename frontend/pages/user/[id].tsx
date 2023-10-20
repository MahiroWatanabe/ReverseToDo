import { useRouter } from "next/router";
import React from "react";
import UserTask from "../../components/UserTask/UserTask";
import SearchUser from "../../components/SearchUser/SearchUser";
import styles from "./user.module.scss";

const UserComponent = () => {
  const router = useRouter();
  const { id } = router.query;

  return (
    <div className={styles["parent-container"]}>
      <UserTask id={id as string} />
      <SearchUser />
    </div>
  );
};

export default UserComponent;

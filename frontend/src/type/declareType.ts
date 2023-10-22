type ReturnUserData = {
  CreatedAt: string;
  DeletedAt: string;
  ID: number;
  Tasks: [];
  UpdatedAt: string;
  email: string;
  username: string;
};

type ReturnTaskData = {
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

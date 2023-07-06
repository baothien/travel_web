import { Navigate, Outlet } from "react-router-dom";
import { toast } from "react-toastify";
import { useEffect } from "react";
import { useAppSelector } from "../hooks";
import { selectUser } from "../modules/Authentication/authSlice";

const PrivateRoute = () => {
  const user = useAppSelector(selectUser);

  useEffect(() => {
    if (!user) {
      toast.info("Bạn không có quyền truy cập vào chức năng này");
    }
  }, []);

  return user ? <Outlet /> : <Navigate to="/" />;
};

export default PrivateRoute;

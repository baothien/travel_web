import React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "../modules/Home/Home";
import DetailPlace from "../modules/DetailPlace/DetailPlace";
import Profile from "../modules/Profile/Profile";
import Login from "../modules/Authentication/pages/Login";
import Register from "../modules/Authentication/pages/Register";
import Map from "../modules/Map/Map";
import CheckIn from "../modules/CheckIn";
import { PostCheckIn } from "../modules/CheckIn/components/PostCheckIn";
import { TestPage } from "../components/TestPage";
import PrivateRoute from "./PrivateRoutes";
import Introduce from "../modules/Introduce";

const AppRoutes: React.FC = () => {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/map" element={<Map />} />
        <Route path="/detail-place/:id" element={<DetailPlace />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        <Route element={<PrivateRoute />}>
          <Route path="profile/*" element={<Profile />} />
          <Route path="/check-in" element={<CheckIn />} />
          <Route path="/post-check-in" element={<PostCheckIn />} />
        </Route>

        <Route path="/introduce" element={<Introduce />} />

        <Route path="/test-page" element={<TestPage />} />
      </Routes>
    </>
  );
};

export default AppRoutes;

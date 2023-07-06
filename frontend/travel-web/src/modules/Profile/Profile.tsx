import React, { useState } from "react";
import "./Profile.scss";

import nullAvatar from "../../assets/img/null-avatar.jpg";

import { Header, Footer } from "../../components";

import { useNavigate, Routes, Route, Link } from "react-router-dom";

import HistoryCheckIn from "./components/HistoryCheckIn/HistoryCheckIn";
import FavoritePlaces from "./components/WishList/FavoritePlaces";

import ChangePassword from "./ChangePassword";
import UserInformation from "./UserInformation";

import { Layout, Menu, Row, Col, Space, Avatar, Typography } from "antd";

import {
  UserOutlined,
  BellOutlined,
  HeartOutlined,
  LockOutlined,
  ExportOutlined,
  EnvironmentOutlined,
} from "@ant-design/icons";

import { authActions, selectUser } from "../Authentication/authSlice";
import { useAppDispatch, useAppSelector } from "../../hooks";
import { NotifyProfile } from "../Notification/NotifyProfile";

const { Content, Sider } = Layout;
const { Title, Text } = Typography;
const { SubMenu } = Menu;

const Profile: React.FC = () => {
  const user = useAppSelector(selectUser);
  const navigate = useNavigate();
  const dispatch = useAppDispatch();

  const handleLogout = () => {
    dispatch(authActions.logoutSuccess());
    navigate("/");
  };

  return (
    <Layout
      style={{
        backgroundColor: "#E5E5E5",
        width: "100%",
        height: "100%",
      }}
      className="primary-font"
    >
      <Header />

      <Content style={{ margin: "6rem 8rem" }} className="">
        <Row
          style={{ marginTop: "0.5rem", marginBottom: "0.2rem" }}
          className="ms-3"
          align="middle"
          gutter={10}
        >
          <Col>
            <Avatar size={45} src={user?.avatar ? user.avatar : nullAvatar} />
          </Col>
          <Col>
            <Space direction="vertical" size={0}>
              <Text className="text-secondary">Tài khoản của</Text>
              <Text strong>
                {user?.full_name ? user.full_name : "Tên người dùng"}
              </Text>
            </Space>
          </Col>

          <Col span={5} offset={3}>
            <Row>
              <Title level={4}>
                {window.location.href.includes("user-information")
                  ? "Thông tin cá nhân"
                  : ""}
                {window.location.href.includes("history-check-in")
                  ? "Lịch sử Check-in"
                  : ""}
                {window.location.href.includes("notification")
                  ? "Thông báo của tôi"
                  : ""}
                {window.location.href.includes("favorite-places")
                  ? "Địa điểm yêu thích"
                  : ""}
                {window.location.href.includes("change-password")
                  ? "Đổi mật khẩu"
                  : ""}
              </Title>
            </Row>
          </Col>
        </Row>

        <Layout
          style={{
            backgroundColor: "#E5E5E5",
          }}
        >
          <Sider style={{}} width={280} className="me-5">
            <Menu
              style={{
                height: "100%",
                borderRight: 0,
                backgroundColor: "#E5E5E5",
                fontWeight: "500",
              }}
              mode="inline"
              defaultSelectedKeys={[
                itemsMenu.filter((item: any) =>
                  window.location.href.includes(item.link)
                )[0].key,
              ]}
            >
              {itemsMenu?.map((item: any) => (
                <Menu.Item key={item.key}>
                  {item.icon}
                  <span>{item.label}</span>
                  <Link to={item.link} />
                </Menu.Item>
              ))}
              <Menu.Item key="6" onClick={handleLogout}>
                <ExportOutlined />
                <span>Đăng xuất</span>
                <Link to="/" />
              </Menu.Item>
            </Menu>
          </Sider>

          <Layout style={{ backgroundColor: "#fff", padding: "1rem 1rem" }}>
            <Content
              className="w-100"
              style={{
                margin: 0,
                minHeight: 280,
              }}
            >
              <Routes>
                <Route path="user-information" element={<UserInformation />} />
                <Route path="history-check-in" element={<HistoryCheckIn />} />
                <Route path="notification" element={<NotifyProfile />} />
                <Route path="favorite-places" element={<FavoritePlaces />} />
                <Route path="change-password" element={<ChangePassword />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      </Content>

      <Footer />
    </Layout>
  );
};

export default Profile;

const itemsMenu: any = [
  {
    key: "1",
    icon: React.createElement(UserOutlined),
    label: "Thông tin cá nhân",
    link: "user-information",
  },
  {
    key: "2",
    icon: React.createElement(EnvironmentOutlined),
    label: "Lịch sử check in",
    link: "history-check-in",
  },
  {
    key: "3",
    icon: React.createElement(BellOutlined),
    label: "Thông báo của tôi",
    link: "notification",
  },
  {
    key: "4",
    icon: React.createElement(HeartOutlined),
    label: "Địa điểm yêu thích",
    link: "favorite-places",
  },
  {
    key: "5",
    icon: React.createElement(LockOutlined),
    label: "Đổi mật khẩu",
    link: "change-password",
  },
];

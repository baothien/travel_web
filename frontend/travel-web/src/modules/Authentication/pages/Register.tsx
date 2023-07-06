import {
  ArrowLeftOutlined,
  LockOutlined,
  MailOutlined,
} from "@ant-design/icons";

import { Link } from "react-router-dom";

import { toast } from "react-toastify";
import { Button, Col, Form, Input, Row, Typography } from "antd";
import React from "react";

import logo from "../../../assets/logo/logo.png";

import "./Authentication.scss";

import apiAccount from "../../../apis/authApi";
import { RegisterParams } from "../../../models/common";

const { Text } = Typography;
const Register: React.FC = () => {
  const onFinish = async (values: any) => {
    const param: RegisterParams = {
      email: values.email,
      password: values.password,
      avatar: "",
      date_of_birth: "",
      full_name: "",
      phone: "",
    };

    await apiAccount
      .postRegister(param)
      .then((res) => {
        toast.success(`${res.message}`);
      })
      .catch((err) => {
        alert(err);
      });
  };

  return (
    <div>
      <div className="text-white primary-font">
        <div className="bg-image"></div>

        <Link to={"/"}>
          <Button
            className="bg-btn d-flex align-items-center justify-content-center"
            type="text"
            icon={<ArrowLeftOutlined className="text-dark" />}
          >
            <Text className="text-dark fs-6">Quay về trang chủ</Text>
          </Button>
        </Link>

        <div className="bg-text text-white rounded">
          {/* <Row>
            <Col>
              <img style={{ width: "50%" }} alt="" src={logo} />
            </Col>
          </Row> */}

          <Row justify="center">
            <Text
              style={{ fontSize: "1.1rem" }}
              className="text-white-50 fw-light mb-4"
            >
              Đăng ký tài khoản cùng Virtual Travel
            </Text>
          </Row>

          <Row className="" justify="center">
            <Form className="login-form" onFinish={onFinish}>
              <Form.Item className="" name="email">
                <Input
                  style={{
                    backgroundColor: "rgba(0,0,0, 0.4)",
                    color: "#fff",
                    padding: "0.5rem",
                    width: "13rem",
                  }}
                  className="border-0 rounded m-0"
                  placeholder="Nhập email"
                  prefix={<MailOutlined />}
                />
              </Form.Item>

              <Form.Item className="" name="password">
                <Input.Password
                  style={{
                    backgroundColor: "rgba(0,0,0, 0.4)",
                    color: "#fff",
                    padding: "0.5rem",
                    width: "13rem",
                  }}
                  className="border-0 rounded"
                  placeholder="Nhập mật khẩu"
                  prefix={<LockOutlined />}
                />
              </Form.Item>

              <Form.Item className="" name="passwordConfirm">
                <Input.Password
                  style={{
                    backgroundColor: "rgba(0,0,0, 0.4)",
                    color: "#fff",
                    padding: "0.5rem",
                    width: "13rem",
                  }}
                  className="border-0 rounded"
                  placeholder="Nhập lại mật khẩu"
                  prefix={<LockOutlined />}
                />
              </Form.Item>

              <Form.Item>
                <Button
                  style={{
                    backgroundColor: "#FD7E14",
                    width: "13rem",
                    height: "2.2rem",
                  }}
                  className="border-0 text-white rounded"
                  htmlType="submit"
                >
                  Đăng ký
                </Button>
              </Form.Item>
            </Form>
          </Row>

          <Row justify="center">
            <Text style={{ fontSize: "0.8rem" }} className="text-white">
              Đã có tài khoản?{" "}
              <Link to={"/login"}>
                <span style={{ color: "#FD7E14" }}> Đăng nhập</span>
              </Link>
            </Text>
          </Row>
        </div>
      </div>
    </div>
  );
};

export default Register;

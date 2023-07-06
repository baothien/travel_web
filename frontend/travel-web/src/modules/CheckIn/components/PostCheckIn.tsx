/* eslint-disable @typescript-eslint/no-empty-interface */
import React, { useState, useEffect } from "react";
import Swal from "sweetalert2";
import NSFWPredictor from "../../../utils/nsfwCheck";
import { CompareSlider } from "./CompareSlider";

import { FloatButton } from "antd";
import { DownloadOutlined, RollbackOutlined } from "@ant-design/icons";

import { Row, Spin } from "antd";
import { useAppDispatch, useAppSelector } from "../../../hooks";
import { selectCapturedImgUrl, selectCurrentPlaceId } from "../checkInSlice";
import { useNavigate } from "react-router-dom";
import uploadApi from "../../../apis/uploadApi";
import checkInApi from "../../../apis/checkInApi";

export interface IPostCheckInProps {}

export function PostCheckIn(props: IPostCheckInProps) {
  const placeId = useAppSelector(selectCurrentPlaceId);

  const [isNsfw, setIsNsfw] = useState(true);

  const [showEnhance, setShowEnhance] = useState(false);

  const navigate = useNavigate();

  const dispatch = useAppDispatch();

  const capturedImgUrl = useAppSelector(selectCapturedImgUrl);

  const [originalPhoto, setOriginalPhoto] = useState<string | null>(null);

  const [restoredImage, setRestoredImage] = useState<string | null>(null);

  const handleSave = (imgUrl) => {
    Swal.fire({
      title: "Đang lưu vào hồ sơ...",
      allowEscapeKey: false,
      allowOutsideClick: false,
      didOpen: () => {
        Swal.showLoading();
      },
    });

    fetch(imgUrl)
      .then((response) => response.blob())
      .then((blob) => {
        const file = new File([blob], "checkedIn.png", { type: blob.type });

        uploadApi
          .upload({ file: file, type: "check_in" })
          .then((res) => {
            const params = {
              name: "checked in",
              place_id: placeId,
              url: res.data.full_path,
            };

            checkInApi
              .postCheckInImg(params)
              .then((res) => {
                Swal.close();

                Swal.fire({
                  icon: "success",
                  title: "Chúc mừng",
                  text: "Ảnh đã được lưu thành công",
                  showDenyButton: true,
                  confirmButtonText: "Tiếp tục check in",
                  denyButtonText: `Lịch sử check in`,
                }).then((result) => {
                  if (result.isConfirmed) {
                    navigate("/check-in");
                  } else if (result.isDenied) {
                    navigate("/profile/history-check-in");
                  }
                });
              })
              .catch((err) => {
                Swal.fire({
                  icon: "error",
                  title: "Không thể lưu ảnh...",
                  allowEscapeKey: false,
                  allowOutsideClick: false,
                  didOpen: () => {
                    Swal.showLoading();
                  },
                });
              });
          })
          .catch((err) => console.log(err));
      });
  };

  useEffect(() => {
    const checkSafe = async () => {
      Swal.fire({
        title: "Đang kiểm tra hình ảnh...",
        allowEscapeKey: false,
        allowOutsideClick: false,
        didOpen: () => {
          Swal.showLoading();
        },
      });

      await fetch(capturedImgUrl)
        .then((response) => response.blob())
        .then(async (blob) => {
          const file = new File([blob], "img-check-in.png", {
            type: blob.type,
          });

          const predict = NSFWPredictor.isSafeImg(file);

          await predict
            .then((res) => {
              if (res) {
                setIsNsfw(false);
                setOriginalPhoto(capturedImgUrl.replace("raw", "thumbnail"));

                Swal.fire({
                  icon: "success",
                  title: "Chúc mừng",
                  text: "Có thể sử dụng ảnh này",
                });
              } else {
                Swal.fire({
                  icon: "error",
                  title: "Ảnh không hợp lệ",
                  text: "Hãy thử check in lại nhé!",
                  showDenyButton: true,
                  confirmButtonText: "Check in",
                  denyButtonText: `Quay về trang chủ`,
                }).then((result) => {
                  if (result.isConfirmed) {
                    navigate("/check-in");
                  } else if (result.isDenied) {
                    navigate("/");
                  }
                });
              }
            })
            .catch((err) => console.log("err", err));
        });
    };

    checkSafe();
  }, []);

  const enhanceImage = async (fileUrl: string) => {
    Swal.fire({
      title: "Đang nâng cấp hình ảnh...",
      allowEscapeKey: false,
      allowOutsideClick: false,
      didOpen: () => {
        Swal.showLoading();
      },
    });

    await new Promise((resolve) => setTimeout(resolve, 500));

    const res = await fetch("https://gog-be.vercel.app/api/generate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ imageUrl: fileUrl }),
    });

    Swal.close();

    if (res.status === 200) {
      const newPhoto = await res.json();
      setRestoredImage(newPhoto);

      Swal.fire({
        icon: "success",
        title: "Chúc mừng",
        text: "Ảnh đã được nâng cấp",
      });

      setShowEnhance(true);
    } else {
      if (res.status === 429) {
        Swal.fire({
          icon: "warning",
          title: "Đã quá lượt sử dụng",
          text: "Hãy thử lại sau 24h",
        });
      } else {
        Swal.fire({
          icon: "error",
          title: "Không thể nâng cấp hình ảnh",
          text: "Hãy thử lại",
        });
      }
    }
  };

  // async function generatePhoto(fileUrl: string) {
  //   await new Promise((resolve) => setTimeout(resolve, 500));

  //   const res = await fetch("https://gog-be.vercel.app/api/generate", {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify({ imageUrl: fileUrl }),
  //   });

  //   if (res.status === 429) {
  //     console.log("Xài quá số lần rùi nè");
  //   } else {
  //     const newPhoto = await res.json();

  //     if (res.status !== 200) {
  //       console.log("Errorrrrrrrrrrrrrrrrrr");
  //     } else {
  //       setRestoredImage(newPhoto);
  //     }
  //   }
  // }

  return (
    <Row style={{ position: "relative" }} justify="center">
      <div className="vw-100 bg-image"></div>

      <Row
        style={{
          position: "absolute",
          top: "5%",
          width: showEnhance ? "55%" : "",
          border: "1px solid #ddd",
          borderRadius: "0.25rem",
          padding: "1rem",
          boxShadow:
            "rgba(0, 0, 0, 0.2) 0px 11px 15px -7px, rgba(0, 0, 0, 0.14) 0px 24px 38px 3px, rgba(0, 0, 0, 0.12) 0px 9px 46px 8px",
        }}
        justify="center"
        align="middle"
      >
        {showEnhance && restoredImage ? (
          <CompareSlider
            original={capturedImgUrl.replace("raw", "thumbnail")}
            restored={restoredImage.replace("raw", "thumbnail")}
          />
        ) : !isNsfw ? (
          <img
            src={capturedImgUrl.replace("raw", "thumbnail")}
            alt="original photo"
            className="rounded"
            width={600}
            height={500}
          />
        ) : (
          <div>
            <Spin size="large" />
          </div>
        )}
      </Row>

      <>
        <FloatButton.Group
          style={{ right: "15%", bottom: "40%" }}
          shape="square"
          type="primary"
        >
          <FloatButton
            tooltip={<div>Check in lại</div>}
            icon={<RollbackOutlined />}
            onClick={() => navigate("/check-in")}
          />

          <FloatButton
            tooltip={<div>Lưu vào hồ sơ</div>}
            onClick={
              showEnhance
                ? () => handleSave(restoredImage)
                : () => handleSave(capturedImgUrl)
            }
            icon={<DownloadOutlined />}
          />

          <FloatButton
            icon={<i className="bi bi-stars"></i>}
            tooltip={<div>Nâng cấp hình ảnh</div>}
            onClick={() =>
              enhanceImage(capturedImgUrl.replace("raw", "thumbnail"))
            }
          />
        </FloatButton.Group>
      </>
    </Row>
  );
}

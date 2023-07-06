/* eslint-disable @typescript-eslint/no-empty-interface */
import {
  Button,
  Divider,
  Modal,
  Pagination,
  PaginationProps,
  Rate,
  Row,
  Space,
  Typography,
  Upload,
} from "antd";

import React, { useState } from "react";
import ReviewContent from "./ReviewContent";
import { useParams } from "react-router-dom";
import { CreateReviewParams, Review } from "../../../../models/review";
import reviewApi from "../../../../apis/reviewApi";
import TextArea from "antd/es/input/TextArea";
import type { RcFile, UploadProps } from "antd/es/upload";
import type { UploadFile } from "antd/es/upload/interface";
import { PlusOutlined } from "@ant-design/icons";
import uploadApi from "../../../../apis/uploadApi";
import { toast } from "react-toastify";
import { useAppSelector } from "../../../../hooks";
import { selectUser } from "../../../Authentication/authSlice";

const { Text } = Typography;

export interface IReviewsContainerProps {
  isCheckIn: boolean;
}

const getBase64 = (file: RcFile): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result as string);
    reader.onerror = (error) => reject(error);
  });

export default function ReviewsContainer({
  isCheckIn,
}: IReviewsContainerProps) {
  const { id } = useParams();

  const user = useAppSelector(selectUser);

  const [loading, setLoading] = useState(false);

  const [listCmt, setListCmt] = React.useState<Review[]>([]);

  const [page, setPage] = React.useState<number>(1);

  const [totalRow, setTotalRow] = React.useState(1);

  const [renderReview, setRenderReview] = useState(false);

  const handleSubmit = () => {
    setLoading(true);

    if (id) {
      const params: CreateReviewParams = {
        description: valueReview,
        place_id: id,
        rate: rate,
        review_img: fileList.map((item) => ({
          name: item.name,
          url: item.url,
        })),
      };

      reviewApi
        .postReview(params)
        .then((res) => {
          setLoading(false);
          toast.success("Đánh giá thành công");
          setRenderReview((prev) => !prev);

          setRate(0);
          setValueReview("");
          setFileList([]);
          // setRenderReview((prev) => !prev);
        })
        .catch((err) => {
          toast.error("Đánh giá thất bại");
        })
        .finally(() => {
          setLoading(false);
        });
    }
  };

  React.useEffect(() => {
    if (id) {
      reviewApi
        .getReviews(id, { page: page, limit: 5 })
        .then((res) => {
          setTotalRow(res.data.total_rows);
          setListCmt(res.data.rows);
        })
        .catch((err) => console.log(err));
    }
  }, [page, renderReview]);

  const onChangePage: PaginationProps["onChange"] = (page) => {
    setPage(page);
  };

  const [previewOpen, setPreviewOpen] = useState(false);
  const [previewImage, setPreviewImage] = useState("");
  const [previewTitle, setPreviewTitle] = useState("");

  const [rate, setRate] = useState(0);
  const [valueReview, setValueReview] = useState("");
  const [fileList, setFileList] = useState<UploadFile[]>([]);

  const uploadButton = (
    <div>
      <PlusOutlined />
      <div style={{ marginTop: 8 }}>Upload</div>
    </div>
  );

  const handlePreview = async (file: UploadFile) => {
    if (!file.url && !file.preview) {
      file.preview = await getBase64(file.originFileObj as RcFile);
    }

    setPreviewImage(file.url || (file.preview as string));
    setPreviewOpen(true);
    setPreviewTitle(
      file.name || file.url!.substring(file.url!.lastIndexOf("/") + 1)
    );
  };

  const handleChangeImage: UploadProps["onChange"] = (
    //   {
    //   fileList: newFileList,
    // }
    value: any
  ) => {
    const fileObj = value.file.originFileObj;
    uploadApi
      .upload({ file: fileObj, type: "review" })
      .then((res) => {
        const resImage = res.data;
        const file: UploadFile = {
          uid: resImage.id,
          name: resImage.name,
          status: "done",
          url: resImage.full_path,
        };
        setFileList((prev) => [...prev, file]);
      })
      .catch((err) => console.log(err));
  };

  const handleCancelUpload = () => setPreviewOpen(false);

  return (
    <div>
      <Row className="w-100" justify="center">
        {user ? (
          isCheckIn ? (
            <Space
              style={{
                width: "68%",
                border: "1px solid #cfcfcf",
                borderRadius: "0.5rem",
                backgroundColor: "#e6e6e6",
              }}
              className="p-2 mb-5"
              direction="vertical"
              size={15}
            >
              <Rate
                className="border"
                allowHalf
                onChange={setRate}
                value={rate}
              />

              <TextArea
                value={valueReview}
                onChange={(e) => setValueReview(e.target.value)}
                placeholder="Viết bình luận"
                autoSize={{ minRows: 2, maxRows: 4 }}
              />

              <Upload
                action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                listType="picture-card"
                fileList={fileList}
                onPreview={handlePreview}
                onChange={handleChangeImage}
              >
                {fileList.length >= 8 ? null : uploadButton}
              </Upload>

              <Space>
                <Button key="submit" type="primary" onClick={handleSubmit}>
                  Xác nhận
                </Button>
              </Space>

              <Modal
                open={previewOpen}
                title={previewTitle}
                footer={null}
                onCancel={handleCancelUpload}
              >
                <img
                  alt="example"
                  style={{ width: "100%" }}
                  src={previewImage}
                />
              </Modal>
            </Space>
          ) : (
            <Space direction="vertical" align="center">
              <Text className="fs-1 text-center" strong>
                Bạn chưa check in địa điểm này
              </Text>

              <Text className="fs-1" strong>
                Hãy <span style={{ color: "#FF7424" }}>check in</span> ngay
              </Text>
            </Space>
          )
        ) : (
          <></>
        )}

        <Divider style={{ borderTop: "#FF7424" }} className="mt-5 mb-5" plain>
          <Text style={{ color: "#FF7424" }} className="fs-5" strong>
            Tất cả các bài đánh giá
          </Text>
        </Divider>

        {listCmt.map((item) => (
          <ReviewContent
            key={item.id}
            cmtId={item.id}
            avatar={item.user.avatar}
            user_email={item.user.email}
            rate={item.rate}
            description={item.description}
            review_img={item.review_img ? item.review_img : []}
            created_at={item.created_at}
            replyList={item.child_review}
          />
        ))}
      </Row>

      <Row className="w-100 mt-3" justify="center">
        <Pagination
          onChange={onChangePage}
          total={totalRow}
          showSizeChanger={false}
          hideOnSinglePage={true}
        />
      </Row>
    </div>
  );
}

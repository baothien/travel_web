import React, { useState } from "react";

import ChildReviewContent from "../ChildReview/ChildReviewContent";

import ButtonLikeReview from "./ButtonLikeReview";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import ReactReadMoreReadLess from "react-read-more-read-less";
import { Typography, Row, Col, Space, Button, Avatar, Image, Rate } from "antd";

import { CommentOutlined, EllipsisOutlined } from "@ant-design/icons";
import { useAppDispatch, useAppSelector } from "../../../../hooks";
import { selectUser } from "../../../Authentication/authSlice";
import { ChildReview } from "../../../../models/review";
import {
  notifyActions,
  selectCmtNotifiedId,
} from "../../../Notification/notifySlice";
import { ChildReviewAction } from "../ChildReview/ChildReviewAction";
import reviewApi from "../../../../apis/reviewApi";

const { Text } = Typography;

export interface ICommentProps {
  cmtId: string;
  avatar: string;
  user_email: string;
  rate: number;
  description: string;
  review_img: any;
  created_at: string;
  replyList: ChildReview[];
}

const ReviewContent: React.FC<ICommentProps> = ({
  cmtId,
  avatar,
  user_email,
  rate,
  description,
  review_img,
  created_at,
}) => {
  const dispatch = useAppDispatch();

  const [replyList, setReplyList] = useState<ChildReview[]>([]);

  const user = useAppSelector(selectUser);

  const [render, reRender] = useState<boolean>(false);

  const [showFullReply, setShowFullReply] = useState(false);

  const [showReplyAction, setShowReplyAction] = useState(false);

  const cmtNotifiedId = useAppSelector(selectCmtNotifiedId);

  React.useEffect(() => {
    const element = document.getElementById(cmtNotifiedId);

    if (element) {
      element.scrollIntoView({ behavior: "smooth", block: "center" });

      dispatch(notifyActions.setCmtNotifiedId(""));
    }
  }, [replyList]);

  React.useEffect(() => {
    reviewApi
      .getReplyReview(cmtId)
      .then((res) => {
        setReplyList(res.data.rows);
      })
      .catch((err) => console.log(err));
  }, [render]);

  return (
    <>
      <Row
        style={{ width: "75%", marginBottom: "5rem" }}
        className=""
        gutter={20}
      >
        <Col span={7}>
          <Row gutter={10} align="middle">
            <Col>
              <Avatar size={50} src={avatar} />
            </Col>
            <Col>
              <Space direction="vertical" size={0}>
                <Text strong>{user_email}</Text>
                <Text className="text-secondary">
                  {created_at.slice(0, 10)}
                </Text>
              </Space>
            </Col>
          </Row>
        </Col>

        <Col span={13}>
          <Row className="mb-1">
            <Rate disabled defaultValue={rate} allowHalf />
          </Row>

          <Row className="mb-1 text-secondary">
            <ReactReadMoreReadLess
              charLimit={190}
              readMoreText={
                <Text style={{ color: "#5449A3" }} strong>
                  Xem thêm
                </Text>
              }
              readLessText={
                <Text style={{ color: "#5449A3" }} strong>
                  Thu gọn
                </Text>
              }
            >
              {description}
            </ReactReadMoreReadLess>
          </Row>

          <Image.PreviewGroup>
            <Space>
              {review_img.map((item: any, index: any) => (
                <Image
                  key={index}
                  style={{ marginRight: "1rem" }}
                  width={120}
                  height={120}
                  src={item.url}
                />
              ))}
            </Space>
          </Image.PreviewGroup>

          {/* reply cmts */}
          <div>
            {showFullReply
              ? replyList.map((item) => (
                  <ChildReviewContent
                    key={item.id}
                    id={item.id}
                    username={item.user.user_name}
                    content={item.description}
                    avatar={item.user.avatar}
                    review_img_url={
                      item.review_img
                        ? item.review_img.map((item) => item.url)
                        : []
                    }
                    createdAt={item.created_at.slice(0, 10)}
                    status={0}
                  />
                ))
              : replyList
                  .slice(0, 4)
                  .map((item) => (
                    <ChildReviewContent
                      key={item.id}
                      id={item.id}
                      username={item.user.user_name}
                      content={item.description}
                      avatar={item.user.avatar}
                      review_img_url={
                        item.review_img
                          ? item.review_img.map((item) => item.url)
                          : []
                      }
                      createdAt={item.created_at.slice(0, 10)}
                      status={0}
                    />
                  ))}
          </div>
          {/* end trả lời */}

          <Row className="w-100 mt-3" justify="center">
            {replyList.length < 5 ? (
              <></>
            ) : showFullReply ? (
              <Text
                style={{
                  color: "#5449A3",
                  cursor: "pointer",
                  fontWeight: "700",
                }}
                onClick={() => setShowFullReply((prev) => !prev)}
              >
                Rút gọn
              </Text>
            ) : (
              <Text
                style={{
                  color: "#CB3837",
                  cursor: "pointer",
                  fontWeight: "600",
                  fontSize: "0.95rem",
                }}
                onClick={() => setShowFullReply((prev) => !prev)}
              >
                Xem thêm
              </Text>
            )}
          </Row>

          {user && showReplyAction && (
            <ChildReviewAction cmtId={cmtId} reRender={reRender} />
          )}
        </Col>

        <Col className="d-flex justify-content-center" span={3}>
          <Space direction="vertical">
            {/* <ButtonLikeReview status={0} /> */}

            <Button
              style={{ width: "4rem", borderRadius: "1rem" }}
              icon={<CommentOutlined />}
              onClick={() => setShowReplyAction((prev) => !prev)}
            />
          </Space>
        </Col>
      </Row>
    </>
  );
};

export default ReviewContent;

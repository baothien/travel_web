import useImageThumbnail from "../hooks/useImageThumbnail";

type ImageButtonProps = {
  imageUrl: string;
  onClick: () => void;
};

function ImageButton(props: ImageButtonProps) {
  const [thumbnailUrl, revokeThumbnailUrl] = useImageThumbnail(props.imageUrl);

  return (
    <button onClick={props.onClick}>
      <img src={thumbnailUrl} alt="" onLoad={revokeThumbnailUrl} />
    </button>
  );
}

export default ImageButton;

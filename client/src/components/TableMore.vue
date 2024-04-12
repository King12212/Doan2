<template>
  <!-- <span class="material-icons"> arrow_drop_up </span> -->
  <div class="drop_down">
    <div class="item" id="item1" @click="viewProfile">
      <img :src="iconview" />
      <div class="text">Hồ sơ sinh viên</div>
    </div>
    <div class="item" id="item2" @click="popup2fa">
      <img :src="icon2fa" />
      <div class="text">Tình trạng điểm danh</div>
    </div>
    <div class="item" id="item3" @click="popupblock">
      <img :src="iconblock" />
      <div class="text">Chặn sinh viên</div>
    </div>
    <div class="item" id="item4" @click="popupdelete">
      <img :src="icondelete" />
      <div class="text">Xóa sinh viên khỏi lớp</div>
    </div>
  </div>
</template>
  
  <script>
import iconview from "../../public/icon/icon_view.png";
import icon2fa from "../../public/icon/icon_2fa.png";
import iconblock from "../../public/icon/icon_block.png";
import icondelete from "../../public/icon/icon_delete.png";
export default {
  name: "TF1",
  setup(_, context) {
    const popup2fa = () => {
      context.emit("view");
    };
    const popupblock = () => {
      context.emit("popup-block");
    };
    const popupdelete = () => {
      context.emit("popup-delete");
    };
    const viewProfile = () => {
      context.emit("view-profile");
    };
    return {
      icon2fa,
      iconblock,
      icondelete,
      iconview,
      popup2fa,
      popupblock,
      popupdelete,
      viewProfile,
    };
  },
};
</script>
  
  <style scoped>
.material-icons {
  position: relative;
  top: 15px;
  color: white;
  z-index: 0;
}
.drop_down {
  width: 250px;
  height: 200px;
  display: flex;
  flex-direction: column;
  border-radius: 5px;
  z-index: 100;
  background-color: white;
  position: fixed;
  z-index: 100;
  margin-top: -230px;
  margin-left: 50px;
  /* position: absolute;
  left: 0px;
  top: 60px; */
  box-shadow: rgba(99, 99, 99, 0.2) 0px 0px 8px 0px;
}
.item {
  height: 40px;
  width: 250px;
  display: flex;
  align-items: center;
  font-size: 14px;
  font-family: "Noto Sans JP";
  font-weight: 500;
  color: #d64d10;
  cursor: pointer;
}
.item:hover {
  background-color: #f0f0f0;
}
#item1 {
  margin-top: 20px;
}
#item4 {
  color: #d00d41;
}
img {
  margin-left: 20px;
}
.text {
  margin-left: 10px;
}
</style>
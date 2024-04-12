<template>
  <div>
    <nav-bar-vue class="nb" />
    <div class="hoso">
      <div class="ava"></div>
      <div class="id">
        <div class="title">Mã số sinh viên</div>
        <div class="content">{{ state.student.Id }}</div>
      </div>
      <div class="ten">
        <div class="title">Tên sinh viên</div>
        <div class="content">
          {{ state.student.Name }}
        </div>
      </div>
      <div class="mail">
        <div class="title">Email</div>
        <div class="content">{{ state.student.Email }}</div>
      </div>
      <div class="phone">
        <div class="title">Số điện thoại</div>
        <div class="content">{{ state.student.Phone }}</div>
      </div>
    </div>
    <div class="diemdanh">
      <div class="box">
        <div class="text">Từ ngày</div>
        <input class="input" type="datetime" />
        <div class="text">Đến ngày</div>
        <input class="input" type="datetime" />
        <button>Nhập</button>
      </div>
      <div class="status">
        <div v-for="week in state.student.List" :key="week.week" class="item">
          <div class="week">{{ week.week }}</div>
          <div class="yes" :class="{ no: !week.status }">
            <span class="material-symbols-outlined"> check_circle </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import NavBarVue from "../components/NavBar.vue";
import { reactive } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
export default {
  components: { NavBarVue },
  setup() {
    const route = useRoute();
    const state = reactive({
      student: {
        Id: route.params.id,
        Name: "Le Huy",
        Email: "huynguyen123@hcmut.edu.vn",
        Phone: "0712098276",
        List: [
          { week: 1, status: false },
          { week: 2, status: false },
          { week: 3, status: false },
          { week: 4, status: false },
          { week: 5, status: false },
          { week: 6, status: true },
          { week: 7, status: false },
          { week: 8, status: false },
          { week: 9, status: false },
          { week: 10, status: false },
        ],
      },
    });
    return { state };
  },
  async created() {
    try {
      const route = useRoute();
      var id = route.params.id;
      const response = await axios.get(
        `http://localhost:8080/getOneStudent/${id}`
      );
      var stt = response.data.student.status;
      for (var i = 0; i < 10; i++) {
        this.state.student.List[i].status = stt[i];
      }
    } catch (e) {
      console.log(e);
    }
  },
  mounted() {
    this.state.student.Name = localStorage.getItem("ten");
    console.log(this.state.student.Name);
  },
};
</script>
<style scoped>
.nb {
  position: fixed;
  top: 0;
  left: 0;
}
.hoso {
  position: fixed;
  top: 150px;
  left: 0;
  width: 500px;
  height: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
  font-size: 18px;
  font-weight: 800;
  /* border: 1px solid lightblue; */
  border-radius: 10px;
  /* background-color:rgb(240,240,240) ; */
  box-shadow: rgba(99, 99, 99, 0.2) 0px 0px 10px 0px;
}
.ava {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  border: 1px solid lightgrey;
}
.id {
  width: 250px;
  display: flex;
  flex-direction: row;
  height: 50px;
  justify-content: space-between;
  align-items: center;
}
.ten {
  width: 250px;
  display: flex;
  flex-direction: row;
  height: 50px;
  justify-content: space-between;
  align-items: center;
}
.mail {
  width: 250px;
  display: flex;
  flex-direction: row;
  height: 50px;
  justify-content: space-between;
  align-items: center;
}
.phone {
  width: 250px;
  display: flex;
  flex-direction: row;
  height: 50px;
  justify-content: space-between;
  align-items: center;
}
.title {
  font-weight: bold;
  margin-left: -100px;
}
.content {
  position: fixed;
  left: 250px;
  color: lightcoral;
}
.diemdanh {
  position: fixed;
  top: 150px;
  left: 600px;
  width: 700px;
  display: flex;
  flex-direction: column;
}
.box {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
.text {
  font-size: 20px;
  font-weight: 700;
}
.input {
  height: 30px;
  padding: 10px;
}
.status {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  margin-top: 100px;
}
.item {
  display: flex;
  flex-direction: column;
  height: 100px;
  width: 100px;
  justify-content: space-between;
  align-items: center;
}
.week {
  font-size: 19px;
  font-weight: 600;
}
.yes {
  color: brown;
}
.no {
  color: grey;
}
</style>
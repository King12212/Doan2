<template>
    <div class = "background">
        <div class = "box">
        <div class ="header"> LOGIN</div>
        <input class = "input" type = "text" v-model="user.username" placeholder="Username" @keyup.enter="login"/>
        <input class = "input" type = "password" v-model="user.password" placeholder="Password" @keyup.enter="login"/>
        <button class ="btn" @click="login">Submit</button>
        
    </div>
    <div class = "warning" v-if="user.right">Username or password is wrong!</div>
    </div>
    
</template>
<script>
import {reactive} from 'vue'
export default{
    name:"LoginView",
    setup(){
        const user = reactive({
            username:"",
            password:"",
            right: false,
        })
        return {user}
    },
    mounted(){
        if (localStorage.getItem('user') === 'true'){
            this.$router.push('/');
        }
    },
    methods:{
        login(){
            if (this.user.username !== "admin" || this.user.password !== "admin"){
                this.user.right = true
            }else{
                localStorage.setItem("user","true")
                this.user.right = false
                this.$router.push("/");
            }
        }
    }
}
</script>
<style scoped>
.background{
    background-color: aquamarine;
    background:linear-gradient(170deg,#201a39 0 50%,#272450 50% 100%);
    position:fixed;
    top: 0;
    left: 0;
    height: 1000px;
    width: 2000px;
    display:flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}
.warning{
    position:fixed;
    color:red;
    top:600px;
    left:810px;
    z-index:100;
    font-size:20px
}
.header{
    margin-top:20px;
    color:rgb(255,128,0);
    font-size:40px;
}
.box{
    display:flex;
    flex-direction: column;
    width :500px;
    height:400px;
    background-color: white;
    align-items: center;
    justify-content: space-around;
    border-radius: 5%;
}
.input{
    position: relative;
    width:400px;
    height:50px;
    font-size:20px; 
    border: solid lightgrey 1px;
    border-radius: 20px;
    padding-left: 10px;
}
input:focus{
    outline:none
}

.btn{
    width:100px;
    height:40px;
    background-color:aquamarine;
    border: solid lightgrey 1px;
    border-radius: 40px;
    cursor:pointer;
}
.btn:hover{
    background-color:aquamarine;
    opacity: 50%;
}
</style>
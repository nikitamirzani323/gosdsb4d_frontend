<script>
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, onValue } from "firebase/database";
    import dayjs from "dayjs";
    const firebaseConfig = {
        apiKey: "AIzaSyCnjwV66P7jDLx5A0Hlh7CHKoZ2tg9jmMY",
        authDomain: "united-rope-233010.firebaseapp.com",
        databaseURL: "https://united-rope-233010-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "united-rope-233010",
        storageBucket: "united-rope-233010.appspot.com",
        messagingSenderId: "994050756260",
        appId: "1:994050756260:web:4dee40c4ca0c34a1842031"
    };


    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);
    const sdsb4dday = ref(db, 'sdsb4dday');
    const sdsb4dnight = ref(db, 'sdsb4dnight');
    let listsdsbday = [];
    let listsdsbnight = [];
    let size_image = "30"
    let size_clock = "50"
    let day_date_draw = ""
    let day_next_draw = ""
    let day_prize1 = ""
    let day_prize2 = ""
    let day_prize3 = ""
    let day_img_1_prize1 = ""
    let day_img_2_prize1 = ""
    let day_img_3_prize1 = ""
    let day_img_4_prize1 = ""
    let day_img_1_prize2 = ""
    let day_img_2_prize2 = ""
    let day_img_3_prize2 = ""
    let day_img_4_prize2 = ""
    let day_img_1_prize3 = ""
    let day_img_2_prize3 = ""
    let day_img_3_prize3 = ""
    let day_img_4_prize3 = ""

    let night_date_draw = ""
    let night_next_draw = ""
    let night_prize1 = ""
    let night_prize2 = ""
    let night_prize3 = ""
    let night_img_1_prize1 = ""
    let night_img_2_prize1 = ""
    let night_img_3_prize1 = ""
    let night_img_4_prize1 = ""
    let night_img_1_prize2 = ""
    let night_img_2_prize2 = ""
    let night_img_3_prize2 = ""
    let night_img_4_prize2 = ""
    let night_img_1_prize3 = ""
    let night_img_2_prize3 = ""
    let night_img_3_prize3 = ""
    let night_img_4_prize3 = ""
    let temp_day = ""
    let temp_day_hour = "00"
    let temp_day_minute = "00"
    let temp_day_second = "00"
    onValue(sdsb4dday, (snapshot) => {
        const data = snapshot.val();
        day_date_draw = data['datedraw']
        day_next_draw = dayjs(data['nextdraw']).format("DD-MMM-YYYY")
        
        day_prize1 = data['prize1']
        day_prize2 = data['prize2']
        day_prize3 = data['prize3']
        day_img_1_prize1 = getImage(day_prize1[0])
        day_img_2_prize1 = getImage(day_prize1[1])
        day_img_3_prize1 = getImage(day_prize1[2])
        day_img_4_prize1 = getImage(day_prize1[3])
        day_img_1_prize2 = getImage(day_prize2[0])
        day_img_2_prize2 = getImage(day_prize2[1])
        day_img_3_prize2 = getImage(day_prize2[2])
        day_img_4_prize2 = getImage(day_prize2[3])
        day_img_1_prize3 = getImage(day_prize3[0])
        day_img_2_prize3 = getImage(day_prize3[1])
        day_img_3_prize3 = getImage(day_prize3[2])
        day_img_4_prize3 = getImage(day_prize3[3])

        temp_day = dayjs(data['nextdraw']+" 14:30:00").valueOf()
        setInterval(function() {
            let now = new Date().getTime();
            let distance = temp_day - now;
            let hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
            let minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            temp_day_hour = hours
            temp_day_minute = minutes
            temp_day_second = seconds
            if(hours < 10){
                temp_day_hour = "0"+hours
            }else{
                temp_day_hour = hours
            }
            if(minutes < 10){
                temp_day_minute = "0"+minutes
            }else{
                temp_day_minute = minutes
            }
            if(seconds < 10){
                temp_day_second = "0"+seconds
            }else{
                temp_day_second = seconds
            }
            if (distance < 0) {
                clearInterval();
                temp_day_hour = "00"
                temp_day_minute = "00"
                temp_day_second = "00"
            }
        }, 1000)
    });
    let temp_night = ""
    let temp_night_hour = "00"
    let temp_night_minute = "00"
    let temp_night_second = "00"
    onValue(sdsb4dnight, (snapshot) => {
        const data = snapshot.val();
        night_date_draw = data['datedraw']
        night_next_draw = dayjs(data['nextdraw']).format("DD-MMM-YYYY")
        night_prize1 = data['prize1']
        night_prize2 = data['prize2']
        night_prize3 = data['prize3']
        night_img_1_prize1 = getImage(night_prize1[0])
        night_img_2_prize1 = getImage(night_prize1[1])
        night_img_3_prize1 = getImage(night_prize1[2])
        night_img_4_prize1 = getImage(night_prize1[3])
        night_img_1_prize2 = getImage(night_prize2[0])
        night_img_2_prize2 = getImage(night_prize2[1])
        night_img_3_prize2 = getImage(night_prize2[2])
        night_img_4_prize2 = getImage(night_prize2[3])
        night_img_1_prize3 = getImage(night_prize3[0])
        night_img_2_prize3 = getImage(night_prize3[1])
        night_img_3_prize3 = getImage(night_prize3[2])
        night_img_4_prize3 = getImage(night_prize3[3])

        temp_night = dayjs(data['nextdraw']+" 21:30:00").valueOf()
        setInterval(function() {
            let now = new Date().getTime();
            let distance = temp_night - now;
            let hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
            let minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            if(hours < 10){
                temp_night_hour = "0"+hours
            }else{
                temp_night_hour = hours
            }
            if(minutes < 10){
                temp_night_minute = "0"+minutes
            }else{
                temp_night_minute = minutes
            }
            if(seconds < 10){
                temp_night_second = "0"+seconds
            }else{
                temp_night_second = seconds
            }
            if (distance < 0) {
            clearInterval();
                temp_night_hour = "00"
                temp_night_minute = "00"
                temp_night_second = "00"
            }
        }, 1000)
    });
    function getImage(e){
        let urlimg = "";
        switch(e){
            case "0":
                urlimg = "number/ball-0.svg";break;
            case "1":
                urlimg = "number/ball-1.svg";break;
            case "2":
                urlimg = "number/ball-2.svg";break;
            case "3":
                urlimg = "number/ball-3.svg";break;
            case "4":
                urlimg = "number/ball-4.svg";break;
            case "5":
                urlimg = "number/ball-5.svg";break;
            case "6":
                urlimg = "number/ball-6.svg";break;
            case "7":
                urlimg = "number/ball-7.svg";break;
            case "8":
                urlimg = "number/ball-8.svg";break;
            case "9":
                urlimg = "number/ball-9.svg";break;
        }
        return urlimg;
    }
    
    async function initSDSB4DDAY() {
        const resPasar = await fetch("/api/listsdsbday", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
            }),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
            let record = jsonPasar.record;
            if (record != null) {
                for (var i = 0; i < 5; i++) {
                listsdsbday = [
                    ...listsdsbday,{
                        sdsbday_date: record[i]["sdsbday_date"],
                        sdsbday_prize1: record[i]["sdsbday_prize1"],
                        sdsbday_prize2: record[i]["sdsbday_prize2"],
                        sdsbday_prize3: record[i]["sdsbday_prize3"]
                    },
                ];
                }
            } else {
                alert("Error");
            }
            } else {
            alert("Error");
            }
        }
    }
    async function initSDSB4DNIGHT() {
        const resPasar = await fetch("/api/listsdsbnight", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
            }),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
            let record = jsonPasar.record;
            if (record != null) {
                for (var i = 0; i < 5; i++) {
                listsdsbnight = [
                    ...listsdsbnight,{
                        sdsbnight_date: record[i]["sdsbnight_date"],
                        sdsbnight_prize1: record[i]["sdsbnight_prize1"],
                        sdsbnight_prize2: record[i]["sdsbnight_prize2"],
                        sdsbnight_prize3: record[i]["sdsbnight_prize3"]
                    },
                ];
                }
            } else {
                alert("Error");
            }
            } else {
            alert("Error");
            }
        }
    }
    initSDSB4DDAY()
    initSDSB4DNIGHT()
</script>
<center>
    <button
        type="button" class="btn btn-info btn-sm" style="background-color: #e91e65;border-color:#e91e65;">Login</button>
    <button
        type="button" class="btn btn-info btn-sm" style="background-color: #e91e65;border-color:#e91e65;">Sign-up</button>
</center>
<div class="card" style="padding:0px;margin:0px;background-color:none;border:none;margin-top:10px;">
    <div class="card-header" style="padding: 0px;margin:0px;background-color:none;">
        <table class="table" style="margin-bottom: -5px;">
            <thead>
                <tr style="background-color: #4ec3dd;border-style: none;border-bottom-color: #4ec3dd;">
                    <th colspan="2" style="text-align: center;vertical-align: top;font-size: 18px;color:#545254;">NEXT DRAWING</th>
                </tr>
                <tr style="background-color: #4ec3dd;">
                    <th width="*" style="text-align: left;vertical-align:top;font-size:12px;color:#545254;">TIME</th>
                    <th NOWRAP style="text-align: right;vertical-align:top;font-size:12px;color:#545254;">{day_next_draw}, 02.30PM</th>
                </tr>
            </thead>
        </table>
      </div>
    <div class="card-body" style="margin: 0px;padding:5px;background-color: #4ec3dd;">
        <div class="container">
            <div class="row align-items-center">
              <div class="col">
                <span class="countdownBox">
                    {temp_day_hour}
                </span>
              </div>
              <div class="col">
                <span class="countdownBox">
                    {temp_day_minute}
                </span>
              </div>
              <div class="col">
                <span class="countdownBox">
                    {temp_day_second}
                </span>
              </div>
            </div>
        </div>
    </div>
</div>
<table class="table" style="margin-top: 5px;">
    <thead>
        <tr style="background-color: #4ec3dd;border-style: none;border-bottom-color: #4ec3dd;">
            <th colspan="4" style="text-align: center;vertical-align: top;font-size: 18px;color:#545254;">SDSB4D DAY</th>
        </tr>
        <tr style="background-color: #4ec3dd;border-bottom-color: #4ec3dd;">
            <th width="*" style="text-align: left;vertical-align:top;font-size:13px;color:#545254;">CURRENT DRAW</th>
            <th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:13px;color:#545254;">{day_date_draw}, 02.30PM</th>
        </tr>
    </thead>
    <tbody style="border-top:none;border-bottom-color: #4ec3dd;">
        <tr style="background-color: #4ec3dd;border-bottom-color: #4ec3dd;">
            <td style="text-align: center;vertical-align:middle;font-size: 15px;font-weight: bold;">1st Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{day_img_1_prize1}" alt="">
                <img width="{size_image}" src="{day_img_2_prize1}" alt="">
                <img width="{size_image}" src="{day_img_3_prize1}" alt="">
                <img width="{size_image}" src="{day_img_4_prize1}" alt="">
            </td>
        </tr>
        <tr style="background-color: #4ec3dd;border-bottom-color: #4ec3dd;">
            <td style="text-align: center;vertical-align:top;font-size: 15px;font-weight: bold;">2nd Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{day_img_1_prize2}" alt="">
                <img width="{size_image}" src="{day_img_2_prize2}" alt="">
                <img width="{size_image}" src="{day_img_3_prize2}" alt="">
                <img width="{size_image}" src="{day_img_4_prize2}" alt="">
            </td>
        </tr>
        <tr style="background-color: #4ec3dd;border-bottom-color: #4ec3dd;">
            <td style="text-align: center;vertical-align:top;font-size: 15px;font-weight: bold;">3rd Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{day_img_1_prize3}" alt="">
                <img width="{size_image}" src="{day_img_2_prize3}" alt="">
                <img width="{size_image}" src="{day_img_3_prize3}" alt="">
                <img width="{size_image}" src="{day_img_4_prize3}" alt="">
            </td>
        </tr>
    </tbody>
</table>
<div class="card" style="padding:0px;margin:0px;background-color:none;border:none;">
    <div class="card-header" style="padding: 0px;margin:0px;background-color:none;">
        <table class="table" style="margin-bottom: -5px;">
            <thead>
                <tr style="background-color: #e4ad68;border-style: none;border-bottom-color: #e4ad68;">
                    <th colspan="4" style="text-align: center;vertical-align: top;font-size: 18px;color:#545254;">NEXT DRAWING</th>
                </tr>
                <tr style="background-color: #e4ad68;">
                    <th width="*" style="text-align: left;vertical-align:top;font-size:12px;color:#545254;">TIME</th>
                    <th colspan="2" NOWRAP style="text-align: right;vertical-align:top;font-size:12px;color:#545254;">{night_next_draw}, 21.30PM</th>
                </tr>
            </thead>
        </table>
      </div>
    <div class="card-body" style="margin: 0px;padding:5px;background-color:#e4ad68;">
        <div class="container">
            <div class="row align-items-center">
              <div class="col">
                <span class="countdownBox2">
                    {temp_night_hour}
                </span>
              </div>
              <div class="col">
                <span class="countdownBox2">
                    {temp_night_minute}
                </span>
              </div>
              <div class="col">
                <span class="countdownBox2">
                    {temp_night_second}
                </span>
              </div>
            </div>
        </div>
    </div>
</div>
<table class="table" style="margin-top: 5px;">
    <thead>
        <tr style="background-color: #e4ad68;border-style: none;border-bottom-color: #e4ad68;">
            <th colspan="4" style="text-align: center;vertical-align: top;font-size: 18px;color:#545254;">SDSB4D NIGHT</th>
        </tr>
        <tr style="background-color: #e4ad68;border-style: none;border-bottom-color: #e4ad68;">
            <th width="*" style="text-align: left;vertical-align:top;font-size:13px;color:#545254;">CURRENT DRAW</th>
            <th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:13px;color:#545254;">{night_date_draw}, 21.30pm</th>
        </tr>
    </thead>
    <tbody style="border-top:none;">
        <tr style="background-color: #e4ad68;border-bottom-color: #e4ad68;">
            <td style="text-align: center;vertical-align:middle;font-size: 15px;font-weight: bold;color:#545254;">1st Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{night_img_1_prize1}" alt="">
                <img width="{size_image}" src="{night_img_2_prize1}" alt="">
                <img width="{size_image}" src="{night_img_3_prize1}" alt="">
                <img width="{size_image}" src="{night_img_4_prize1}" alt="">
            </td>
        </tr>
        <tr style="background-color: #e4ad68;border-bottom-color: #e4ad68;">
            <td style="text-align: center;vertical-align:top;font-size: 15px;font-weight: bold;color:#545254;">2nd Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{night_img_1_prize2}" alt="">
                <img width="{size_image}" src="{night_img_2_prize2}" alt="">
                <img width="{size_image}" src="{night_img_3_prize2}" alt="">
                <img width="{size_image}" src="{night_img_4_prize2}" alt="">
            </td>
        </tr>
        <tr style="background-color: #e4ad68;border-bottom-color: #e4ad68;">
            <td style="text-align: center;vertical-align:top;font-size: 15px;font-weight: bold;color:#545254;">3rd Prize</td>
            <td style="text-align: center;vertical-align:top;">
                <img width="{size_image}" src="{night_img_1_prize3}" alt="">
                <img width="{size_image}" src="{night_img_2_prize3}" alt="">
                <img width="{size_image}" src="{night_img_3_prize3}" alt="">
                <img width="{size_image}" src="{night_img_4_prize3}" alt="">
            </td>
        </tr>
    </tbody>
</table>
<table class="table">
    <thead>
        <tr style="background-color: #4ec3dd;border-style: none;border-bottom-color: #4ec3dd;">
            <th colspan="4" style="text-align: center;vertical-align: top;color:#545254;">LATEST RESULT</th>
        </tr>
        <tr style="background-color: #4ec3dd;">
            <th width="*" style="text-align: center;vertical-align:top;font-size: 13px;color:#545254;">DATE</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 1</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 2</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 3</th>
        </tr>
    </thead>
    <tbody style="border-top:none;">
        {#each listsdsbday as rec }
        <tr style="background-color: #f2f7f5;">
            <td style="text-align: center;vertical-align:top;font-size: 13px;">{rec.sdsbday_date}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;">{rec.sdsbday_prize1}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;">{rec.sdsbday_prize2}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;">{rec.sdsbday_prize3}</td>
        </tr>
        {/each}
    </tbody>
</table>
<table class="table">
    <thead>
        <tr style="background-color: #e4ad68;border-style: none;border-bottom-color: #e4ad68;">
            <th colspan="4" style="text-align: center;vertical-align: top;color:#545254;">LATEST RESULT</th>
        </tr>
        <tr style="background-color: #e4ad68;">
            <th width="*" style="text-align: center;vertical-align:top;font-size: 13px;color:#545254;">DATE</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 1</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 2</th>
            <th width="25%" style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">PRIZE 3</th>
        </tr>
    </thead>
    <tbody style="border-top:none;">
        {#each listsdsbnight as rec }
        <tr style="background-color: #fcf9fb;">
            <td style="text-align: center;vertical-align:top;font-size: 13px;color:#545254;">{rec.sdsbnight_date}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">{rec.sdsbnight_prize1}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">{rec.sdsbnight_prize2}</td>
            <td style="text-align: right;vertical-align:top;font-size: 13px;color:#545254;">{rec.sdsbnight_prize3}</td>
        </tr>
        {/each}
    </tbody>
</table>
<style>
    .countdownBox {
        border: 3px #4ec3dd solid;
        color: #545254;
        padding: 5px;
        text-align: center;
        display: block;
        font-size: 30px;
        box-shadow: 0px 3px 1px -2px rgb(0 0 0 / 20%), 0px 2px 2px 0px rgb(0 0 0 / 14%), 0px 1px 5px 0px rgb(0 0 0 / 12%);
    }
    .countdownBox2 {
        border: 3px #e4ad68 solid;
        color: #545254;
        padding: 5px;
        text-align: center;
        display: block;
        font-size: 30px;
        box-shadow: 0px 3px 1px -2px rgb(0 0 0 / 20%), 0px 2px 2px 0px rgb(0 0 0 / 14%), 0px 1px 5px 0px rgb(0 0 0 / 12%);
    }
</style>
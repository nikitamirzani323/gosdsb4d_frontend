<script>
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, onValue } from "firebase/database";
    
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
    let size_image = "40"
    let day_date_draw = ""
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
    onValue(sdsb4dday, (snapshot) => {
        const data = snapshot.val();
        day_date_draw = data['datedraw']
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
    });
    onValue(sdsb4dnight, (snapshot) => {
        const data = snapshot.val();
        night_date_draw = data['datedraw']
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
    
</script>
<div class="row">
    <div class="col-sm-6">
        <table class="table">
            <thead>
                <tr style="background-color: #3e266d;border-style: none;border-bottom-color: #3e266d;">
                    <th colspan="4" style="text-align: center;vertical-align: top;font-size: 30px;color:white;">SDSB4D DAY</th>
                </tr>
                <tr style="background-color: #3e266d;">
                    <th width="*" style="text-align: left;vertical-align:top;font-size:15px;color:white;">CURRENT DRAW</th>
                    <th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:15px;color:white;">{day_date_draw}, 02.30PM</th>
                </tr>
            </thead>
            <tbody style="border-top:none;">
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:middle;font-size: 18px;font-weight: bold;">1st Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
                        <img width="{size_image}" src="{day_img_1_prize1}" alt="">
                        <img width="{size_image}" src="{day_img_2_prize1}" alt="">
                        <img width="{size_image}" src="{day_img_3_prize1}" alt="">
                        <img width="{size_image}" src="{day_img_4_prize1}" alt="">
                    </td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">2nd Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
                        <img width="{size_image}" src="{day_img_1_prize2}" alt="">
                        <img width="{size_image}" src="{day_img_2_prize2}" alt="">
                        <img width="{size_image}" src="{day_img_3_prize2}" alt="">
                        <img width="{size_image}" src="{day_img_4_prize2}" alt="">
                    </td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">3rd Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
                        <img width="{size_image}" src="{day_img_1_prize3}" alt="">
                        <img width="{size_image}" src="{day_img_2_prize3}" alt="">
                        <img width="{size_image}" src="{day_img_3_prize3}" alt="">
                        <img width="{size_image}" src="{day_img_4_prize3}" alt="">
                    </td>
                </tr>
            </tbody>
        </table>
        <table class="table">
            <thead>
                <tr style="background-color: #3e266d;border-style: none;border-bottom-color: #3e266d;">
                    <th colspan="4" style="text-align: center;vertical-align: top;color:white;">LATEST RESULT</th>
                </tr>
                <tr style="background-color: #3e266d;">
                    <th width="*" style="text-align: center;vertical-align:top;font-size: 15px;color:white;">DATE</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;color:white;">PRIZE 1</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;color:white;">PRIZE 2</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;color:white;">PRIZE 3</th>
                </tr>
            </thead>
            <tbody style="border-top:none;">
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
            </tbody>
        </table>
        <div class="d-grid">
            <button type="button" class="btn btn-warning" style="background-color: #3e266d;border-color: #3e266d;">ALL RESULT</button>
        </div>
    </div>
    <div class="col-sm-6">
        <table class="table">
            <thead>
                <tr style="background-color: #f9da7c;border-style: none;border-bottom-color: #f9da7c;">
                    <th colspan="4" style="text-align: center;vertical-align: top;font-size: 30px">SDSB4D NIGHT</th>
                </tr>
                <tr style="background-color: #f9da7c;border-style: none;border-bottom-color: #f9da7c;">
                    <th width="*" style="text-align: left;vertical-align:top;font-size:15px;">CURRENT DRAW</th>
                    <th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:15px;">{night_date_draw}, 21.30pm</th>
                </tr>
            </thead>
            <tbody style="border-top:none;">
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:middle;font-size: 18px;font-weight: bold;">1st Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
                        <img width="{size_image}" src="{night_img_1_prize1}" alt="">
                        <img width="{size_image}" src="{night_img_2_prize1}" alt="">
                        <img width="{size_image}" src="{night_img_3_prize1}" alt="">
                        <img width="{size_image}" src="{night_img_4_prize1}" alt="">
                    </td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">2nd Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
                        <img width="{size_image}" src="{night_img_1_prize2}" alt="">
                        <img width="{size_image}" src="{night_img_2_prize2}" alt="">
                        <img width="{size_image}" src="{night_img_3_prize2}" alt="">
                        <img width="{size_image}" src="{night_img_4_prize2}" alt="">
                    </td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">3rd Prize</td>
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">
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
                <tr style="background-color: #f9da7c;border-style: none;border-bottom-color: #f9da7c;">
                    <th colspan="4" style="text-align: center;vertical-align: top;">LATEST RESULT</th>
                </tr>
                <tr style="background-color: #f9da7c;">
                    <th width="*" style="text-align: center;vertical-align:top;font-size: 15px;">DATE</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 1</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 2</th>
                    <th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 3</th>
                </tr>
            </thead>
            <tbody style="border-top:none;">
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
                <tr style="background-color: #f2f7f5;">
                    <td style="text-align: center;vertical-align:top;font-size: 15px;">2021-10-27</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">1234</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">4321</td>
                    <td style="text-align: right;vertical-align:top;font-size: 15px;">9821</td>
                </tr>
            </tbody>
        </table>
        <div class="d-grid">
            <button type="button" class="btn btn-danger" style="background-color: #f9da7c;border-color:#f9da7c;color:black;">ALL RESULT</button>
        </div>
    </div>
</div>
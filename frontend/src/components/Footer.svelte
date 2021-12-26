<script>
	export let client_device = "WEBSITE";
	let font_size = "";
	let img_size = "";
	let img_size_visa = "";
	let img_size_mastercard = "";
	if(client_device == "WEBSITE"){
		font_size = "15px";
	}else{
		font_size = "13px";
		img_size = "70";
		img_size_visa = "70";
		img_size_mastercard = "";
	}
</script>
<footer class="footer mt-auto py-3" style="background-color: #ffd401;">
	<div class="container">
		<center>
			<span style="color:#020c13;font-size:{font_size}">We accept these payment method</span><br>
			<img width="{img_size_visa}" src="visa.svg" alt="">
			<img width="{img_size_mastercard}" src="mastercard.svg" alt="">
			<img width="{img_size}" src="paypal.svg" alt="">
		</center>
		<br>
		<div class="row">
			<div class="col-sm-9">
				<center style="color:#020c13;font-size:{font_size}">
					<p>
						Players must be 18 years old. If you do not have control, gambling may be harmful. Please, play with responsibility.
						All efforts will be made to ensure accuracy of the prizes, winnings and other information posted on sdsb4d.com. Official Winning Numbers are those that are selected in the respective drawings and are controlled by an independent accounting firm.
						In case of non-compliance, official winning results will prevail.
						The "SDSB4D" trademark registration number 958786573-gf-94515 is owned and operated by the Lotera Association of SDBD4D Lotera.
					</p>
				</center>
			</div>
			<div class="col-sm-3">
				<center>
					<img src="google-play.png" alt="">
				</center>
			</div>
		</div>
		<div class="clearfix"></div> <br>
		<center style="color:#020c13;font-size:{font_size}">
			Copyright © sdsb4d.com
		</center>
	</div>
</footer>
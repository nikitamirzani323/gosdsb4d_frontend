<script>
	import Navbar from "./components/Navbar.svelte"
	import Footer from "./components/Footer.svelte"
	
	export let name;

	
	let css_loader = "display: none;";
    let msgloader = "";
</script>

<Navbar />
<div class="container" style="margin-top: 70px;">
	<div class="row">
		<div class="col-sm-6">
			<table class="table">
				<thead>
					<tr style="background-color: #8bd3dd;">
						<th colspan="4" style="text-align: center;vertical-align: top;font-size: 30px">SDSB4D DAY</th>
					</tr>
					<tr style="background-color: #8bd3dd;">
						<th width="*" style="text-align: left;vertical-align:top;font-size:15px;">CURRENT DRAW</th>
						<th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:15px;">27-Oct-2021, 01.00pm</th>
					</tr>
				</thead>
				<tbody>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:middle;font-size: 18px;font-weight: bold;">1st Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-0.svg" alt="">
							<img width="50" src="number/ball-1.svg" alt="">
							<img width="50" src="number/ball-2.svg" alt="">
							<img width="50" src="number/ball-3.svg" alt="">
						</td>
					</tr>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">2nd Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-4.svg" alt="">
							<img width="50" src="number/ball-5.svg" alt="">
							<img width="50" src="number/ball-6.svg" alt="">
							<img width="50" src="number/ball-7.svg" alt="">
						</td>
					</tr>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">3rd Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-8.svg" alt="">
							<img width="50" src="number/ball-9.svg" alt="">
							<img width="50" src="number/ball-0.svg" alt="">
							<img width="50" src="number/ball-1.svg" alt="">
						</td>
					</tr>
				</tbody>
			</table>
			<table class="table">
				<thead>
					<tr style="background-color: #8bd3dd;">
						<th colspan="4" style="text-align: center;vertical-align: top;">LATEST RESULT</th>
					</tr>
					<tr style="background-color: #8bd3dd;">
						<th width="*" style="text-align: center;vertical-align:top;font-size: 15px;">DATE</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 1</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 2</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 3</th>
					</tr>
				</thead>
				<tbody>
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
				<button type="button" class="btn btn-warning" style="background-color: #8bd3dd;border-color: #8bd3dd;">ALL RESULT</button>
			</div>
		</div>
		<div class="col-sm-6">
			<table class="table">
				<thead>
					<tr style="background-color: #fe98a3;">
						<th colspan="4" style="text-align: center;vertical-align: top;font-size: 30px">SDSB4D NIGHT</th>
					</tr>
					<tr style="background-color: #fe98a3;">
						<th width="*" style="text-align: left;vertical-align:top;font-size:15px;">CURRENT DRAW</th>
						<th NOWRAP width="50%" style="text-align: right;vertical-align:top;font-size:15px;">27-Oct-2021, 01.00pm</th>
					</tr>
				</thead>
				<tbody>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:middle;font-size: 18px;font-weight: bold;">1st Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-0.svg" alt="">
							<img width="50" src="number/ball-1.svg" alt="">
							<img width="50" src="number/ball-2.svg" alt="">
							<img width="50" src="number/ball-3.svg" alt="">
						</td>
					</tr>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">2nd Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-4.svg" alt="">
							<img width="50" src="number/ball-5.svg" alt="">
							<img width="50" src="number/ball-6.svg" alt="">
							<img width="50" src="number/ball-7.svg" alt="">
						</td>
					</tr>
					<tr style="background-color: #f2f7f5;">
						<td style="text-align: center;vertical-align:top;font-size: 18px;font-weight: bold;">3rd Prize</td>
						<td style="text-align: center;vertical-align:top;font-size: 15px;">
							<img width="50" src="number/ball-8.svg" alt="">
							<img width="50" src="number/ball-9.svg" alt="">
							<img width="50" src="number/ball-0.svg" alt="">
							<img width="50" src="number/ball-1.svg" alt="">
						</td>
					</tr>
				</tbody>
			</table>
			<table class="table">
				<thead>
					<tr style="background-color: #fe98a3;">
						<th colspan="4" style="text-align: center;vertical-align: top;">LATEST RESULT</th>
					</tr>
					<tr style="background-color: #fe98a3;">
						<th width="*" style="text-align: center;vertical-align:top;font-size: 15px;">DATE</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 1</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 2</th>
						<th width="25%" style="text-align: right;vertical-align:top;font-size: 15px;">PRIZE 3</th>
					</tr>
				</thead>
				<tbody>
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
				<button type="button" class="btn btn-danger">ALL RESULT</button>
			</div>
		</div>
	</div>
</div>
<div class="clear-fix"></div><br>
<Footer />

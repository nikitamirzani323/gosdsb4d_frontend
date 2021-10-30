<script>
	
	import Navbar from "./components/Navbar.svelte"
	import Footer from "./components/Footer.svelte"
	import Home from "./pages/Home.svelte"
	
	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}

	
</script>

<Navbar />
<div class="contentmodif" style="margin-top: -10px;">
	<div class="clearfix"></div><br>
	<div class="container content">
		<Home />
	</div>
	<div class="clearfix"></div><br>
</div>
<Footer />

<style>
	.container {
		width: 1000px;
		max-width: none !important;
	}
	.content{
		background-color: rgba(44, 31, 58, 0.5);
	}
	.contentmodif{
    	background-image: url('https://i.ibb.co/vZpPWtJ/Lottery-balls-isolated-on-orange-background.jpg');
        background-position:center center;
        background-repeat:no-repeat;
	}
</style>
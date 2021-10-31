<script>
	
	import Navbar from "./components/Navbar.svelte"
	import Navbar_mobile from "./components/Navbar_mobile.svelte"
	import Footer from "./components/Footer.svelte"
	import Home from "./pages/Home.svelte"
	import Home_mobile from "./pages/Home_mobile.svelte"
	
	let client_device = "";
	if(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
    	// true for mobile device
    	client_device = "MOBILE";
	}else{
    	client_device = "WEBSITE";
	}

	console.log(client_device)
</script>

{#if client_device == "WEBSITE"}
<Navbar />
<div class="contentmodif" style="margin-top: -10px;">
	<div class="clearfix"></div><br>
	<div class="container content">
		<Home />
	</div>
	<div class="clearfix"></div><br>
</div>
<div class="clearfix"></div>
<br>
<Footer {client_device} />
<style>
	.container {
		width: 1000px;
		max-width: none !important;
	}
	.content{
		margin-top:30px;
		border-radius:10px;
		background-color: rgba(233, 154, 0, 0.5);
	}
	.contentmodif{
    	background-image: url('background.jpg');
        background-position:center center;
        background-repeat:no-repeat;
	}
</style>
{:else}
<Navbar_mobile />
<div class="container" style="margin-top: 60px;">
	<Home_mobile />
</div>
<Footer {client_device} />
{/if}
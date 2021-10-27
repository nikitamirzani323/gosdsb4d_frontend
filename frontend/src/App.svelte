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

	
	let css_loader = "display: none;";
    let msgloader = "";
</script>

<Navbar />
<div class="container" style="margin-top: 40px;">
	<Home />
</div>
<div class="clear-fix"></div><br>
<Footer />

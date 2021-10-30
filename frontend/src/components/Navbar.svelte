<script>
    import Modal from "../components/Modal.svelte"
    let myModal = "";

    const handlePopup = () => {
        myModal = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal.show();
    };
</script>
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container">
      <a class="navbar-brand" href="#">
        <img style="margin-top:-10px;" src="logo.png" alt="" />
      </a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="#">Home</a>
          </li>
          <li class="nav-item">
            <a 
                on:click={() => {
                    handlePopup();
                }} 
                class="nav-link" href="#">Claim Prize</a>
          </li>
          <li class="nav-item">
            <a 
                on:click={() => {
                    handlePopup();
                }} 
                class="nav-link" href="#">About</a>
          </li>
        </ul>
      </div>
      
      <form class="d-flex">
        <button
            on:click={() => {
                handlePopup();
            }}  
            type="button" class="btn btn-info me-2" style="background-color: #e91e65;border-color:#e91e65;">Login</button>
        <button
            on:click={() => {
                handlePopup();
            }} 
            type="button" class="btn btn-info" style="background-color: #e91e65;border-color:#e91e65;">Sign-up</button>
      </form>
    </div>
</nav>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title=""
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <center>
            Authentication Required
		    <br>
		    Users must be located in Europe
        </center>
	</slot:template>
</Modal>
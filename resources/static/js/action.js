function closeAlert() {
    $(".alert").delay(5000).slideUp(200, function () {
        $(this).alert('close');
    });
}

function ok(json) {
    var payload = "<div class='alert alert-success'>"
    payload += text;
    payload += "</div>";
    return payload
}

function fail(text) {
    var payload = "<div class='alert alert-danger'>"
    payload += text;
    payload += "</div>";
    return payload
}

$("#submitPing").submit(function (e) {
    e.preventDefault();
    var form = $("#submitPing");
    var addr = form.find("select[name='ip']").val();
    var req = $.get("/ewallet/action", { method: "ping", addr: ip });
    req.done(function (data) {
        if (data.status === "ok") {
            $("#notifPing").append(ok("Sukses melakukan ping ke " + addr + "."));
            closeAlert();
        } else {
            $("#notifPing").append(fail("Tidak dapat melakukan ping ke " + addr + "."));
            closeAlert();
        }
    });
});

$("#submitRegister").submit(function (e) {
    e.preventDefault();
    var form = $("#submitRegister");
    var id = form.find("input[name='id']").val();
    var nama = form.find("input[name='nama']").val();
    var addr = form.find("select[name='ip']").val();
    var req = $.get("/ewallet/action", { method: "register", id: id, nama: nama, addr: ip });
    req.done(function (data) {
        if (data.status === "ok") {
            $("#notifRegister").append(ok("Sukses register user " + nama + " dengan id " + id + "."));
            closeAlert();
        } else {
            $("#notifRegister").append(fail("Gagal melakukan register. " + data.reason));
            closeAlert();
        }
    });
});
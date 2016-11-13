function closeAlert() {
    $(".alert").delay(5000).slideUp(200, function () {
        $(this).alert('close');
    });
}

function ok(text) {
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
    var req = $.get("https://prakash.sisdis.ui.ac.id/ewallet/action?method=ping&addr=" + addr);
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
    var req = $.get("https://prakash.sisdis.ui.ac.id/ewallet/action?method=register&id=" + id + "&nama=" + nama + "&addr=" + addr);
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

$("#submitGetSaldo").submit(function (e) {
    e.preventDefault();
    var form = $("#submitGetSaldo");
    var id = form.find("input[name='id']").val();
    var req = $.get("https://prakash.sisdis.ui.ac.id/ewallet/action?method=getSaldo&id=" + id);
    req.done(function (data) {
        if (data.status === "ok") {
            $("#notifGetSaldo").append(ok("Jumlah saldo untuk user dengan id " + id + " adalah " + data.nilai_saldo + "."));
            closeAlert();
        } else {
            $("#notifGetSaldo").append(fail("Gagal mengecek saldo. " + data.reason));
            closeAlert();
        }
    });
});

$("#submitGetTotalSaldo").submit(function (e) {
    e.preventDefault();
    var form = $("#submitGetTotalSaldo");
    var id = form.find("input[name='id']").val();
    var req = $.get("https://prakash.sisdis.ui.ac.id/ewallet/action?method=getTotalSaldo&id=" + id);
    req.done(function (data) {
        if (data.status === "ok") {
            $("#notifGetTotalSaldo").append(ok("Jumlah total saldo untuk user dengan id " + id + " adalah " + data.nilai_saldo + "."));
            closeAlert();
        } else {
            $("#notifGetTotalSaldo").append(fail("Gagal mengecek total saldo. " + data.reason));
            closeAlert();
        }
    });
});

$("#submitTransfer").submit(function (e) {
    e.preventDefault();
    var form = $("#submitTransfer");
    var id = form.find("input[name='id']").val();
    var jumlah = form.find("input[name='jumlah']").val();
    var addr = form.find("select[name='ip']").val();
    var req = $.get("https://prakash.sisdis.ui.ac.id/ewallet/action?method=transfer&id=" + id + "&nilai=" + jumlah + "&addr=" + addr);
    req.done(function (data) {
        if (data.status === "ok") {
            $("#notifTransfer").append(ok("Berhasil melakukan transfer untuk user dengan id " + id + " ke cabang " + addr + " dengan jumlah " + jumlah + "."));
            closeAlert();
        } else {
            $("#notifTransfer").append(fail("Gagal melakukan transfer. " + data.reason));
            closeAlert();
        }
    });
});
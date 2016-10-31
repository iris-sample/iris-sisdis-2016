package main

import "github.com/kataras/iris"

func ewalletWSDL(c *iris.Context) {
	var data = `<wsdl:definitions xmlns="http://schemas.xmlsoap.org/wsdl/"
xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/"
xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/"
xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/"
xmlns:xsd="http://www.w3.org/2001/XMLSchema"
xmlns:tns="id.ac.ui.cs.sisdis.KantorCabangNamespace"
targetNamespace="id.ac.ui.cs.sisdis.KantorCabangNamespace">
  <wsdl:types>
    <xsd:schema targetNamespace="id.ac.ui.cs.sisdis.KantorCabangNamespace" />
  </wsdl:types>
  <message name="getSaldoRequest">
    <part name="user_id" type="xsd:string" />
  </message>
  <message name="getSaldoResponse">
    <part name="getSaldoReturn" type="xsd:int" />
  </message>
  <message name="getTotalSaldoRequest">
    <part name="user_id" type="xsd:string" />
  </message>
  <message name="getTotalSaldoResponse">
    <part name="getTotalSaldoReturn" type="xsd:int" />
  </message>
  <message name="pingRequest" />
  <message name="pingResponse">
    <part name="pingReturn" type="xsd:int" />
  </message>
  <message name="registerRequest">
    <part name="user_id" type="xsd:string" />
    <part name="nama" type="xsd:string" />
    <part name="ip_domisili" type="xsd:string" />
  </message>
  <message name="transferRequest">
    <part name="user_id" type="xsd:string" />
    <part name="nilai" type="xsd:string" />
  </message>
  <message name="transferResponse">
    <part name="transferReturn" type="xsd:int" />
  </message>
  <wsdl:portType name="KantorCabangPortType">
    <wsdl:operation name="getSaldo">
      <wsdl:input message="tns:getSaldoRequest" />
      <wsdl:output message="tns:getSaldoResponse" />
    </wsdl:operation>
    <wsdl:operation name="getTotalSaldo">
      <wsdl:input message="tns:getTotalSaldoRequest" />
      <wsdl:output message="tns:getTotalSaldoResponse" />
    </wsdl:operation>
    <wsdl:operation name="ping">
      <wsdl:input message="tns:pingRequest" />
      <wsdl:output message="tns:pingResponse" />
    </wsdl:operation>
    <wsdl:operation name="register">
      <wsdl:input message="tns:registerRequest" />
    </wsdl:operation>
    <wsdl:operation name="transfer">
      <wsdl:input message="tns:transferRequest" />
      <wsdl:output message="tns:transferResponse" />
    </wsdl:operation>
  </wsdl:portType>
  <binding name="KantorCabangBinding"
  type="tns:KantorCabangPortType">
    <soap:binding style="rpc"
    transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="getSaldo">
      <soap:operation soapAction="KantorCabang#getSaldo"
      style="rpc" />
      <wsdl:input>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="getTotalSaldo">
      <soap:operation soapAction="KantorCabang#getTotalSaldo"
      style="rpc" />
      <wsdl:input>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="ping">
      <soap:operation soapAction="KantorCabang#ping" style="rpc" />
      <wsdl:input>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="register">
      <soap:operation soapAction="KantorCabang#register"
      style="rpc" />
      <wsdl:input>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:input>
    </wsdl:operation>
    <wsdl:operation name="transfer">
      <soap:operation soapAction="KantorCabang#transfer"
      style="rpc" />
      <wsdl:input>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="encoded"
        namespace="id.ac.ui.cs.sisdis.KantorCabangNamespace"
        encodingStyle="http://schemas.xmlsoap.org/soap/encoding/" />
      </wsdl:output>
    </wsdl:operation>
  </binding>
  <wsdl:service name="KantorCabang">
    <wsdl:port name="KantorCabangPort"
    binding="tns:KantorCabangBinding">
      <soap:address location="http://localhost:7070/ewallet/server" />
    </wsdl:port>
  </wsdl:service>
</wsdl:definitions>`
	c.SetContentType("text/xml")
	c.SetBodyString(data)
}

// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package bookv1beta1

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func NewShipmentTrackingEventServiceCommand(config aipcli.Config) *cobra.Command {
	return aipcli.NewServiceCommand(
		config,
		File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto.
			Services().ByName("ShipmentTrackingEventService"),
		map[protoreflect.FullName]string{
			"einride.saga.extend.book.v1beta1.ShipmentTrackingEventService": " Shipment tracking event service.\n",
		},
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto.
				Services().ByName("ShipmentTrackingEventService").Methods().ByName("CreateShipmentTrackingEvent"),
			&CreateShipmentTrackingEventRequest{},
			&ShipmentTrackingEvent{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.Address.city":                                               " City\n",
				"einride.saga.extend.book.v1beta1.Address.co":                                                 " Care of (C/O)\n",
				"einride.saga.extend.book.v1beta1.Address.contact_info":                                       " Contact information\n",
				"einride.saga.extend.book.v1beta1.Address.display_name":                                       " The displayed name of the address\n",
				"einride.saga.extend.book.v1beta1.Address.lat_lng":                                            " The geographic location of the address\n",
				"einride.saga.extend.book.v1beta1.Address.line1":                                              " Address line 1\n",
				"einride.saga.extend.book.v1beta1.Address.line2":                                              " Address line 2\n",
				"einride.saga.extend.book.v1beta1.Address.postal_code":                                        " Postal code\n",
				"einride.saga.extend.book.v1beta1.Address.recipient":                                          " Recipient\n",
				"einride.saga.extend.book.v1beta1.Address.region_code":                                        " Region code (Unicode CLDR region code)\n https://cldr.unicode.org/\n",
				"einride.saga.extend.book.v1beta1.Address.state_code":                                         " State code\n",
				"einride.saga.extend.book.v1beta1.CreateShipmentTrackingEventRequest.parent":                  " The parent shipment in which to create the event.\n",
				"einride.saga.extend.book.v1beta1.CreateShipmentTrackingEventRequest.shipment_tracking_event": " The shipment tracking event\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.address":                              " The address where the event happened.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.code":                                 " The kind of tracking event that happened\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.create_time":                          " The creation timestamp of the ShipmentTrackingEvent. When this event was received by Einride.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.delete_time":                          " The deletion timestamp of the ShipmentTrackingEvent. Set if the event is deleted.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.description":                          " Description for the ShipmentTrackingEvent.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.name":                                 " The resource name of the ShipmentTrackingEvent.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.occurred_time":                        " When the event occurred.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.shipment":                             " Resource name of the shipment that owns the ShipmentTrackingEvent.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.update_time":                          " The last update timestamp of the ShipmentTrackingEvent.\n Updated when create/update/delete operation is performed.\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEvent.vehicle":                              " On what vehicle did this event occur\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEventService.CreateShipmentTrackingEvent":   " Create a tracking event for a shipment\n\n This is an AIP standard [Create](https://google.aip.dev/133) method.\n",
				"einride.saga.extend.book.v1beta1.Vehicle.carrier_reference_id":                               " Carrier reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.driver_reference_id":                                " Driver reference id\n",
				"einride.saga.extend.book.v1beta1.Vehicle.vehicle_reference_id":                               " Vehicle reference id\n",
				"google.protobuf.Timestamp.nanos":                                                             " Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.\n",
				"google.protobuf.Timestamp.seconds":                                                           " Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.\n",
				"google.type.LatLng.latitude":                                                                 " The latitude in degrees. It must be in the range [-90.0, +90.0].\n",
				"google.type.LatLng.longitude":                                                                " The longitude in degrees. It must be in the range [-180.0, +180.0].\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto.
				Services().ByName("ShipmentTrackingEventService").Methods().ByName("GetShipmentTrackingEvent"),
			&GetShipmentTrackingEventRequest{},
			&ShipmentTrackingEvent{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.GetShipmentTrackingEventRequest.name":                  " The resource name of the shipment tracking event to retrieve.\n Format: spaces/{space}/shipments/{shipment}/trackingEvents/{shipmentTrackingEvent}\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEventService.GetShipmentTrackingEvent": " Get a shipment tracking event.\n\n This is an AIP standard [Get](https://google.aip.dev/131) method.\n",
			},
		),
		aipcli.NewMethodCommand(
			config,
			File_einride_saga_extend_book_v1beta1_shipment_tracking_event_service_proto.
				Services().ByName("ShipmentTrackingEventService").Methods().ByName("ListShipmentTrackingEvents"),
			&ListShipmentTrackingEventsRequest{},
			&ListShipmentTrackingEventsResponse{},
			map[protoreflect.FullName]string{
				"einride.saga.extend.book.v1beta1.ListShipmentTrackingEventsRequest.page_size":             " Requested page size. Server may return fewer shipment tracking events than requested.\n If unspecified, server will pick an appropriate default.\n",
				"einride.saga.extend.book.v1beta1.ListShipmentTrackingEventsRequest.page_token":            " A token identifying a page of results the server should return.\n Typically, this is the value of\n [ListShipmentTrackingEventsResponse.page_token][einride.saga.extend.book.v1beta1.ListShipmentTrackingEventsResponse.page_token]\n returned from the previous call to `ListShipmentTrackingEvents` method.\n",
				"einride.saga.extend.book.v1beta1.ListShipmentTrackingEventsRequest.parent":                " The resource name of the parent, which owns this collection of shipment tracking events.\n Format: spaces/{space}/shipments/{shipment}/trackingEvents\n",
				"einride.saga.extend.book.v1beta1.ShipmentTrackingEventService.ListShipmentTrackingEvents": " List shipment tracking events for a shipment.\n The events are ordered by ascending occurred_time.\n\n See: https://google.aip.dev/132 (Standard methods: List).\n",
			},
		),
	)
}

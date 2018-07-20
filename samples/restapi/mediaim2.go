package main

type Im2MediaFile struct {
	Infomedia struct {
		PhysicalScreens struct {
			PhysicalScreen []struct {
				Name   string `json:"_Name"`
				Type   string `json:"_Type"`
				Width  string `json:"_Width"`
				Height string `json:"_Height"`
				ID     string `json:"_Id,omitempty"`
			} `json:"PhysicalScreen"`
		} `json:"PhysicalScreens"`
		VirtualDisplays struct {
			VirtualDisplay []struct {
				Name         string `json:"_Name"`
				CyclePackage string `json:"_CyclePackage"`
				Width        string `json:"_Width"`
				Height       string `json:"_Height"`
			} `json:"VirtualDisplay"`
		} `json:"VirtualDisplays"`
		MasterPresentation struct {
			MasterCycles struct {
				MasterCycle struct {
					MasterSection struct {
						Duration string `json:"_Duration"`
						Layout   string `json:"_Layout"`
					} `json:"MasterSection"`
					Name string `json:"_Name"`
				} `json:"MasterCycle"`
			} `json:"MasterCycles"`
			MasterEventCycles string `json:"MasterEventCycles"`
			MasterLayouts     struct {
				MasterLayout struct {
					PhysicalScreen []struct {
						VirtualDisplay struct {
							X      string `json:"_X"`
							Y      string `json:"_Y"`
							ZIndex string `json:"_ZIndex"`
							Ref    string `json:"_Ref"`
						} `json:"VirtualDisplay"`
						Ref string `json:"_Ref"`
					} `json:"PhysicalScreen"`
					Name string `json:"_Name"`
				} `json:"MasterLayout"`
			} `json:"MasterLayouts"`
		} `json:"MasterPresentation"`
		Evaluations struct {
			Evaluation []struct {
				And struct {
					StringCompare []struct {
						Generic struct {
							Lang   string `json:"_Lang"`
							Table  string `json:"_Table"`
							Column string `json:"_Column"`
							Row    string `json:"_Row"`
						} `json:"Generic"`
						Value string `json:"_Value"`
					} `json:"StringCompare"`
				} `json:"And,omitempty"`
				Name string `json:"_Name"`
				Or   struct {
					StringCompare []struct {
						Generic struct {
							Lang   string `json:"_Lang"`
							Table  string `json:"_Table"`
							Column string `json:"_Column"`
							Row    string `json:"_Row"`
						} `json:"Generic"`
						Value string `json:"_Value"`
					} `json:"StringCompare"`
				} `json:"Or,omitempty"`
				IntegerCompare struct {
					Generic struct {
						Lang   string `json:"_Lang"`
						Table  string `json:"_Table"`
						Column string `json:"_Column"`
						Row    string `json:"_Row"`
					} `json:"Generic"`
					Begin string `json:"_Begin"`
					End   string `json:"_End"`
				} `json:"IntegerCompare,omitempty"`
				StringCompare struct {
					Generic struct {
						Lang   string `json:"_Lang"`
						Table  string `json:"_Table"`
						Column string `json:"_Column"`
						Row    string `json:"_Row"`
					} `json:"Generic"`
					Value string `json:"_Value"`
				} `json:"StringCompare,omitempty"`
				CodeConversion struct {
					FileName string `json:"_FileName"`
					UseImage string `json:"_UseImage"`
				} `json:"CodeConversion,omitempty"`
				Not struct {
					StringCompare struct {
						Generic struct {
							Lang   string `json:"_Lang"`
							Table  string `json:"_Table"`
							Column string `json:"_Column"`
							Row    string `json:"_Row"`
						} `json:"Generic"`
						Value string `json:"_Value"`
					} `json:"StringCompare"`
				} `json:"Not,omitempty"`
				Format struct {
					RegexReplace struct {
						RegexReplace struct {
							RegexReplace struct {
								RegexReplace struct {
									RegexReplace struct {
										RegexReplace struct {
											RegexReplace struct {
												RegexReplace struct {
													RegexReplace struct {
														RegexReplace struct {
															RegexReplace struct {
																RegexReplace struct {
																	RegexReplace struct {
																		RegexReplace struct {
																			RegexReplace struct {
																				RegexReplace struct {
																					RegexReplace struct {
																						RegexReplace struct {
																							RegexReplace struct {
																								RegexReplace struct {
																									RegexReplace struct {
																										RegexReplace struct {
																											RegexReplace struct {
																												RegexReplace struct {
																													RegexReplace struct {
																														RegexReplace struct {
																															RegexReplace struct {
																																Generic struct {
																																	Lang   string `json:"_Lang"`
																																	Table  string `json:"_Table"`
																																	Column string `json:"_Column"`
																																	Row    string `json:"_Row"`
																																} `json:"Generic"`
																																Pattern     string `json:"_Pattern"`
																																Replacement string `json:"_Replacement"`
																															} `json:"RegexReplace"`
																															Pattern     string `json:"_Pattern"`
																															Replacement string `json:"_Replacement"`
																														} `json:"RegexReplace"`
																														Pattern     string `json:"_Pattern"`
																														Replacement string `json:"_Replacement"`
																													} `json:"RegexReplace"`
																													Pattern     string `json:"_Pattern"`
																													Replacement string `json:"_Replacement"`
																												} `json:"RegexReplace"`
																												Pattern     string `json:"_Pattern"`
																												Replacement string `json:"_Replacement"`
																											} `json:"RegexReplace"`
																											Pattern     string `json:"_Pattern"`
																											Replacement string `json:"_Replacement"`
																										} `json:"RegexReplace"`
																										Pattern     string `json:"_Pattern"`
																										Replacement string `json:"_Replacement"`
																									} `json:"RegexReplace"`
																									Pattern     string `json:"_Pattern"`
																									Replacement string `json:"_Replacement"`
																								} `json:"RegexReplace"`
																								Pattern     string `json:"_Pattern"`
																								Replacement string `json:"_Replacement"`
																							} `json:"RegexReplace"`
																							Pattern     string `json:"_Pattern"`
																							Replacement string `json:"_Replacement"`
																						} `json:"RegexReplace"`
																						Pattern     string `json:"_Pattern"`
																						Replacement string `json:"_Replacement"`
																					} `json:"RegexReplace"`
																					Pattern     string `json:"_Pattern"`
																					Replacement string `json:"_Replacement"`
																				} `json:"RegexReplace"`
																				Pattern     string `json:"_Pattern"`
																				Replacement string `json:"_Replacement"`
																			} `json:"RegexReplace"`
																			Pattern     string `json:"_Pattern"`
																			Replacement string `json:"_Replacement"`
																		} `json:"RegexReplace"`
																		Pattern     string `json:"_Pattern"`
																		Replacement string `json:"_Replacement"`
																	} `json:"RegexReplace"`
																	Pattern     string `json:"_Pattern"`
																	Replacement string `json:"_Replacement"`
																} `json:"RegexReplace"`
																Pattern     string `json:"_Pattern"`
																Replacement string `json:"_Replacement"`
															} `json:"RegexReplace"`
															Pattern     string `json:"_Pattern"`
															Replacement string `json:"_Replacement"`
														} `json:"RegexReplace"`
														Pattern     string `json:"_Pattern"`
														Replacement string `json:"_Replacement"`
													} `json:"RegexReplace"`
													Pattern     string `json:"_Pattern"`
													Replacement string `json:"_Replacement"`
												} `json:"RegexReplace"`
												Pattern     string `json:"_Pattern"`
												Replacement string `json:"_Replacement"`
											} `json:"RegexReplace"`
											Pattern     string `json:"_Pattern"`
											Replacement string `json:"_Replacement"`
										} `json:"RegexReplace"`
										Pattern     string `json:"_Pattern"`
										Replacement string `json:"_Replacement"`
									} `json:"RegexReplace"`
									Pattern     string `json:"_Pattern"`
									Replacement string `json:"_Replacement"`
								} `json:"RegexReplace"`
								Pattern     string `json:"_Pattern"`
								Replacement string `json:"_Replacement"`
							} `json:"RegexReplace"`
							Pattern     string `json:"_Pattern"`
							Replacement string `json:"_Replacement"`
						} `json:"RegexReplace"`
						Pattern     string `json:"_Pattern"`
						Replacement string `json:"_Replacement"`
					} `json:"RegexReplace"`
					Format string `json:"_Format"`
				} `json:"Format,omitempty"`
				Generic struct {
					Lang   string `json:"_Lang"`
					Table  string `json:"_Table"`
					Column string `json:"_Column"`
					Row    string `json:"_Row"`
				} `json:"Generic,omitempty"`
			} `json:"Evaluation"`
		} `json:"Evaluations"`
		Cycles struct {
			StandardCycles struct {
				StandardCycle []struct {
					StandardSection struct {
						Duration string `json:"_Duration"`
						Layout   string `json:"_Layout"`
					} `json:"StandardSection"`
					Name    string `json:"_Name"`
					Enabled struct {
						Or struct {
							Not struct {
								Evaluation struct {
									Ref string `json:"_Ref"`
								} `json:"Evaluation"`
							} `json:"Not"`
							Evaluation []struct {
								Ref string `json:"_Ref"`
							} `json:"Evaluation"`
						} `json:"Or"`
					} `json:"Enabled,omitempty"`
				} `json:"StandardCycle"`
			} `json:"StandardCycles"`
			EventCycles struct {
				EventCycle []struct {
					Enabled struct {
						Evaluation struct {
							Ref string `json:"_Ref"`
						} `json:"Evaluation"`
					} `json:"Enabled,omitempty"`
					Trigger struct {
						Generic struct {
							Lang   string `json:"_Lang"`
							Table  string `json:"_Table"`
							Column string `json:"_Column"`
							Row    string `json:"_Row"`
						} `json:"Generic"`
					} `json:"Trigger"`
					StandardSection struct {
						Duration string `json:"_Duration"`
						Layout   string `json:"_Layout"`
					} `json:"StandardSection"`
					Name string `json:"_Name"`
				} `json:"EventCycle"`
			} `json:"EventCycles"`
		} `json:"Cycles"`
		CyclePackages struct {
			CyclePackage []struct {
				StandardCycles struct {
					StandardCycle []struct {
						Ref string `json:"_Ref"`
					} `json:"StandardCycle"`
				} `json:"StandardCycles"`
				EventCycles struct {
					EventCycle struct {
						Ref string `json:"_Ref"`
					} `json:"EventCycle"`
				} `json:"EventCycles"`
				Name string `json:"_Name"`
			} `json:"CyclePackage"`
		} `json:"CyclePackages"`
		Pools   string `json:"Pools"`
		Layouts struct {
			Layout []struct {
				Resolution []struct {
					Text []struct {
						Font struct {
							Face         string `json:"_Face"`
							Height       string `json:"_Height"`
							Weight       string `json:"_Weight"`
							Italic       string `json:"_Italic"`
							Color        string `json:"_Color"`
							OutlineColor string `json:"_OutlineColor"`
						} `json:"Font"`
						LastPosition struct {
							Location struct {
								X string `json:"X"`
								Y string `json:"Y"`
							} `json:"Location"`
							Size struct {
								Width  string `json:"Width"`
								Height string `json:"Height"`
							} `json:"Size"`
							X      string `json:"X"`
							Y      string `json:"Y"`
							Width  string `json:"Width"`
							Height string `json:"Height"`
						} `json:"LastPosition"`
						X           string `json:"_X"`
						Y           string `json:"_Y"`
						Width       string `json:"_Width"`
						Height      string `json:"_Height"`
						ZIndex      string `json:"_ZIndex"`
						Align       string `json:"_Align"`
						VAlign      string `json:"_VAlign,omitempty"`
						Overflow    string `json:"_Overflow"`
						ScrollSpeed string `json:"_ScrollSpeed"`
						Value       string `json:"_Value,omitempty"`
						Visible     struct {
							Evaluation struct {
								Ref string `json:"_Ref"`
							} `json:"Evaluation"`
						} `json:"Visible,omitempty"`
						Value struct {
							Generic struct {
								Lang   string `json:"_Lang"`
								Table  string `json:"_Table"`
								Column string `json:"_Column"`
								Row    string `json:"_Row"`
							} `json:"Generic"`
						} `json:"Value,omitempty"`
					} `json:"Text,omitempty"`
					Image []struct {
						Visible struct {
							Generic struct {
								Lang   string `json:"_Lang"`
								Table  string `json:"_Table"`
								Column string `json:"_Column"`
								Row    string `json:"_Row"`
							} `json:"Generic"`
						} `json:"Visible,omitempty"`
						X        string `json:"_X"`
						Y        string `json:"_Y"`
						Width    string `json:"_Width"`
						Height   string `json:"_Height"`
						ZIndex   string `json:"_ZIndex"`
						Filename string `json:"_Filename"`
					} `json:"Image,omitempty"`
					Width  string `json:"_Width"`
					Height string `json:"_Height"`
				} `json:"Resolution"`
				Name string `json:"_Name"`
			} `json:"Layout"`
		} `json:"Layouts"`
		Fonts struct {
			Font struct {
				Path string `json:"_Path"`
			} `json:"Font"`
		} `json:"Fonts"`
		XmlnsXsi string `json:"_xmlns:xsi"`
		XmlnsXsd string `json:"_xmlns:xsd"`
		Version  string `json:"_Version"`
		Created  string `json:"_Created"`
	} `json:"Infomedia"`
}
